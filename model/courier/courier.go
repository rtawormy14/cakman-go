package courier

import (
	"bytes"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/go-sql-driver/mysql"

	"github.com/rtawormy14/cakman-go/util/database"
)

// Courierer interface
type Courierer interface {
	GetCourier(int64) (Courier, error)
	AuthenticateCourier(string, string) (Courier, error)
	Insert(Courier, *sqlx.Tx) error
	Update(Courier, *sqlx.Tx) error
	UpdateLocation(Courier, *sqlx.Tx) error
	Remove(Courier, *sqlx.Tx) error
}

// Courier struct
type Courier struct {
	ID            int64     `db:"id" json:"id"`
	Username      string    `db:"username" json:"username"`
	Password      string    `db:"password" json:"-"`
	Name          string    `db:"name" json:"name"`
	Phone         string    `db:"phone" json:"phone"`
	Email         string    `db:"email" json:"email"`
	CreateBy      int64     `db:"create_by" json:"create_by"`
	CreateTime    time.Time `db:"create_time" json:"create_time"`
	UpdateBy      int64     `db:"update_by" json:"update_by"`
	UpdateTime    time.Time `db:"update_time" json:"update_time"`
	Lattitude     float64   `db:"lattitude" json:"lattitude"`
	Longitude     float64   `db:"longitude" json:"longitude"`
	LocUpdateTime time.Time `db:"update_position_time" json:"update_position_time"`
}

// NewCourier is ...
func NewCourier() Courierer {
	return &Courier{}
}

// GetCourier will return courier object if exist
func (c *Courier) GetCourier(code int64) (courier Courier, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return courier, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM courier WHERE id = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&courier, query, code)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Courier][GetCourier] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// AuthenticateCourier is...
func (c *Courier) AuthenticateCourier(username string, password string) (courier Courier, err error) {
	db := database.DB
	//validate parameter
	if username == "" {
		return courier, errors.New("username is mandatory")
	}
	if password == "" {
		return courier, errors.New("password is mandatory")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT * FROM courier WHERE username = ? AND password = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&courier, query, username, password)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Courier][AuthenticateCourier] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// Insert will insert courier data
func (c *Courier) Insert(courier Courier, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "INSERT INTO courier (username,password,name,phone,email,create_by,create_time) VALUES (?,?,?,?,?,?,?)"
	query = tx.Rebind(query)
	tx.MustExec(query, courier.Username, courier.Password, courier.Name, courier.Phone, courier.Email, courier.CreateBy, courier.CreateTime)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Courier][Insert] Error when commit to db : %s \n data : %v ", err, courier)
			err = errors.New("error when insert data to courier")
		}
	}
	return err
}

// UpdateLocation is ....
func (c *Courier) UpdateLocation(courier Courier, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "UPDATE courier SET longitude=?, lattitude=?, update_position_time=? WHERE id=?"
	query = tx.Rebind(query)
	tx.MustExec(query, courier.Longitude, courier.Lattitude, courier.LocUpdateTime, courier.ID)
	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Courier][UpdateLocation] Error when commit to db : %s \n data : %v ", err, courier)
			err = errors.New("error when update location courier")
		}
	}
	return err
}

// Update will update courier data
func (c *Courier) Update(courier Courier, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "UPDATE courier SET username=?, password=?, name=?, phone=?, email=?, update_by=?, update_time=? WHERE id=?"
	query = tx.Rebind(query)
	tx.MustExec(query, courier.Username, courier.Password, courier.Name, courier.Phone, courier.Email, courier.UpdateBy, courier.UpdateTime, courier.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Courier][Update] Error when commit to db : %s \n data : %v ", err, courier)
			err = errors.New("error when update data to courier")
		}
	}
	return err
}

// Remove will remove city data
func (c *Courier) Remove(courier Courier, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "DELETE FROM courier WHERE id=?"
	query = tx.Rebind(query)
	tx.MustExec(query, courier.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Courier][Delete] Error when commit to db : %s \n data : %v ", err, courier)
			err = errors.New("error when update data to courier")
		}
	}
	return err
}
