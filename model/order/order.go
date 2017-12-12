package order

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/go-sql-driver/mysql"

	cityModel "github.com/rtawormy14/cakman-go/model/city"
	countryModel "github.com/rtawormy14/cakman-go/model/country"
	provinceModel "github.com/rtawormy14/cakman-go/model/province"
	"github.com/rtawormy14/cakman-go/util/database"
)

// Orderer interface
type Orderer interface {
	FindResi(string) (Order, error)
	GetOrder(int64) (Order, error)
	GetOrderList(int64, int64, Order) ([]Order, error)
	Insert(Order, *sqlx.Tx) error
	Update(Order, *sqlx.Tx) error
	Remove(Order, *sqlx.Tx) error
}

// Order struct
type Order struct {
	ID         int64     `db:"id" json:"id"`
	Resi       string    `db:"resi" json:"resi"`
	Consignee  string    `db:"consignee" json:"consignee"`
	Phone      string    `db:"phone" json:"phone"`
	Status     string    `db:"status" json:"status"`
	Address    string    `db:"address" json:"address"`
	CountryID  int64     `db:"country_code" json:"-"`
	ProvinceID int64     `db:"province_code" json:"-"`
	CityID     int64     `db:"city_code" json:"-"`
	CreateBy   int64     `db:"create_by" json:"create_by"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateBy   int64     `db:"update_by" json:"update_by"`
	UpdateTime time.Time `db:"update_time" json:"update_time"`

	Country  countryModel.Country   `json:"country,omitempty"`
	Province provinceModel.Province `json:"province,omitempty"`
	City     cityModel.City         `json:"city,omitempty"`
}

const (
	// StatusWarehouse means order is not yet picked up by courier, and it still in warehouse
	StatusWarehouse = "WAREHOUSE"
	// StatusInprogress means order is already picked up by courier and on his way to delivering order
	StatusInprogress = "INPROGRESS"
	// StatusFinish means order is already delivered successfully
	StatusFinish = "FINISH"
)

// NewOrder is ...
func NewOrder() Orderer {
	return &Order{}
}

// FindResi will return Order object if exist
func (o *Order) FindResi(resi string) (order Order, err error) {
	db := database.DB

	//validate parameter
	if resi == "" {
		return order, errors.New("resi is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM orders WHERE resi = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&order, query, resi)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Order][FindResi] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetOrder will return Order object if exist
func (o *Order) GetOrder(orderID int64) (order Order, err error) {
	db := database.DB

	//validate parameter
	if orderID <= 0 {
		return order, errors.New("orderID is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM orders WHERE id = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&order, query, orderID)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Order][GetOrder] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	return
}

// GetOrderList is...
func (o *Order) GetOrderList(page int64, limit int64, filter Order) (orders []Order, err error) {
	db := database.DB

	orders = make([]Order, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM orders ")
	queryBuffer.WriteString("WHERE status = ? ")
	queryBuffer.WriteString("ORDER BY create_time ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET ? LIMIT ?")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&orders, query, StatusWarehouse, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&orders, query, StatusWarehouse)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Order][GetOrderList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	return
}

// Insert will insert Order data
func (o *Order) Insert(order Order, tx *sqlx.Tx) (err error) {
	//TODO
	return
}

// Update will update Order data
func (o *Order) Update(order Order, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := `UPDATE orders 
				SET 
					consignee=?, 
					phone=?, 
					status=?, 
					address=?, 
					country_code=?, 
					province_code=?,  
					city_code=?,
					update_by=?, 
					update_time=?
				WHERE id=?`
	query = tx.Rebind(query)
	tx.MustExec(query, order.Consignee, order.Phone, order.Status, order.Address, order.CountryID, order.ProvinceID, order.CityID, order.UpdateBy, time.Now(), order.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Order][Update] Error when commit to db : %s \n data : %v ", err, order)
			err = errors.New("error when insert data to orders")
		}
	}
	return err
}

// Remove will remove order data
func (o *Order) Remove(order Order, tx *sqlx.Tx) (err error) {
	//TODO
	return
}
