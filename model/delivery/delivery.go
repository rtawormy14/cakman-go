package delivery

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/lib/pq"

	courierModel "github.com/rtawormy14/cakman-go/model/courier"
	orderModel "github.com/rtawormy14/cakman-go/model/order"
	"github.com/rtawormy14/cakman-go/util/database"
)

// Deliverier interface
type Deliverier interface {
	GetDelivery(int64) (Delivery, error)
	GetDeliveryByOrderID(int64) (Delivery, error)
	GetDeliveryHistory(int64, int64, int64) ([]Delivery, error)
	GetDeliveryList(int64, int64, Delivery) ([]Delivery, error)
	Insert(Delivery, *sqlx.Tx) (Delivery, error)
	Update(Delivery, *sqlx.Tx) error
	Remove(Delivery, *sqlx.Tx) error
}

// Delivery struct
type Delivery struct {
	ID         int64     `db:"id" json:"id"`
	OrderID    int64     `db:"order_id" json:"-"`
	CourierID  int64     `db:"courier_id" json:"-"`
	Status     string    `db:"status" json:"status"`
	Note       string    `db:"note" json:"note"`
	CreateBy   int64     `db:"create_by" json:"create_by"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateBy   int64     `db:"update_by" json:"update_by"`
	UpdateTime time.Time `db:"update_time" json:"update_time"`

	Courier courierModel.Courier `json:"courier,omitempty"`
	Order   orderModel.Order     `json:"order,omitempty"`
	History []History            `json:"history,omitempty"`
}

var (
	// StatusPickup means Delivery is in progress
	StatusPickup = "PICKUP"
	// StatusCancel means Delivery is cancelled by courier due to some reasons
	StatusCancel = "CANCEL"
	// StatusFinish means Delivery is already delivered successfully
	StatusFinish = "FINISH"

	// ValidStatus is to used to assert if status delivery is valid or not
	ValidStatus = map[string]bool{
		"PICKUP": true,
		"CANCEL": true,
		"FINISH": true,
	}
)

// NewDelivery is ...
func NewDelivery() Deliverier {
	return &Delivery{}
}

// GetDelivery will return Delivery object if exist
func (d *Delivery) GetDelivery(deliveryID int64) (delivery Delivery, err error) {
	db := database.DB

	//validate parameter
	if deliveryID <= 0 {
		return delivery, errors.New("deliveryID is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM delivery WHERE id = $1 LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&delivery, query, deliveryID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Delivery][GetDelivery] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetDeliveryByOrderID will return Delivery object if exist
func (d *Delivery) GetDeliveryByOrderID(orderID int64) (delivery Delivery, err error) {
	db := database.DB

	//validate parameter
	if orderID <= 0 {
		return delivery, errors.New("deliveryID is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM delivery WHERE order_id = $1 AND status IN($2,$3) LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&delivery, query, orderID, StatusPickup, StatusFinish)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Delivery][GetDelivery] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetDeliveryHistory will return Delivery object if exist
func (d *Delivery) GetDeliveryHistory(page int64, limit int64, courierID int64) (deliveries []Delivery, err error) {
	db := database.DB

	deliveries = make([]Delivery, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM delivery ")
	queryBuffer.WriteString("WHERE courier_id = $1 AND status IN($2,$3) ")
	queryBuffer.WriteString("ORDER BY update_time DESC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET $4 LIMIT $5")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&deliveries, query, courierID, StatusCancel, StatusFinish, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&deliveries, query, courierID, StatusCancel, StatusFinish)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Delivery][GetDeliveryList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	return
}

// GetDeliveryList is...
func (d *Delivery) GetDeliveryList(page int64, limit int64, filter Delivery) (deliveries []Delivery, err error) {
	db := database.DB

	deliveries = make([]Delivery, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM delivery ")
	queryBuffer.WriteString("WHERE courier_id = $1 AND status = $2 ")
	queryBuffer.WriteString("ORDER BY create_time ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET $3 LIMIT $4")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&deliveries, query, filter.CourierID, StatusPickup, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&deliveries, query, filter.CourierID, StatusPickup)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Delivery][GetDeliveryList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	return
}

// Insert will insert Delivery data
func (d *Delivery) Insert(delivery Delivery, tx *sqlx.Tx) (deliveryObj Delivery, err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "INSERT INTO delivery (courier_id,order_id,status,note,create_by,create_time) VALUES ($1,$2,$3,$4,$5,$6) returning id"
	query = tx.Rebind(query)

	var id int64
	tx.QueryRow(query, delivery.CourierID, delivery.OrderID, delivery.Status, delivery.Note, delivery.CreateBy, delivery.CreateTime).Scan(&id)
	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Delivery][Insert] Error when commit to db : %s \n data : %v ", err, delivery)
			err = errors.New("error when insert data to delivery")
		}
	}
	delivery.ID = id
	deliveryObj = delivery

	return
}

// Update will update Delivery data
func (d *Delivery) Update(delivery Delivery, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := `UPDATE delivery 
				SET 
					status=$1, 
					note=$2, 
					update_by=$3, 
					update_time=$4 
				WHERE id=$5`
	query = tx.Rebind(query)
	tx.MustExec(query, delivery.Status, delivery.Note, delivery.UpdateBy, time.Now(), delivery.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Delivery][Update] Error when commit to db : %s \n data : %v ", err, delivery)
			err = errors.New("error when insert data to delivery")
		}
	}
	return err
}

// Remove will remove delivery data
func (d *Delivery) Remove(delivery Delivery, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "DELETE FROM delivery WHERE id = $1"
	query = tx.Rebind(query)
	tx.MustExec(query, delivery.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Delivery][Remove] Error when commit to db : %s \n data : %v ", err, delivery)
			err = errors.New("error when remove data to delivery")
		}
	}
	return err
}
