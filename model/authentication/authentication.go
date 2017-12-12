package authentication

import (
	"database/sql"
	//postgre implementation
	"bytes"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/rtawormy14/cakman-go/util/database"
)

// Authenticator interface
type Authenticator interface {
	GetAuthentication(int64) (auth Authentication, err error)
	GetAuthenticationByToken(string) (auth Authentication, err error)
	Update(Authentication, *sqlx.Tx) error
	Insert(Authentication, *sqlx.Tx) error
	Remove(Authentication, *sqlx.Tx) error
}

// Authentication struct
type Authentication struct {
	ID         int64     `db:"id" json:"-"`
	CourierID  int64     `db:"courier_id" json:"courier_id"`
	Token      string    `db:"token" json:"token"`
	ExpireTime time.Time `db:"expire_time" json:"expire_at"`
	ExpireIn   int64     `json:"expire_in"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}

// NewAuthentication is ...
func NewAuthentication() Authenticator {
	return &Authentication{}
}

// GetAuthentication will return authentication data using id as parameter
func (a *Authentication) GetAuthentication(id int64) (auth Authentication, err error) {
	db := database.DB

	//validate parameter
	if id <= 0 {
		return auth, errors.New("id is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT id, courier_id, token, create_time, expire_time FROM session WHERE id = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&auth, query, id)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Authentication][GetAuthentication] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	auth.ExpireIn = int64(time.Since(auth.ExpireTime).Seconds())
	return
}

// GetAuthenticationByToken will return authentication data using token as parameter
func (a *Authentication) GetAuthenticationByToken(token string) (auth Authentication, err error) {
	db := database.DB

	//validate parameter
	if token == "" || len(token) != 32 {
		return auth, errors.New("token is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT id, courier_id, token, create_time, expire_time FROM session WHERE token = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&auth, query, token)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Authentication][GetAuthenticationByToken] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}

	auth.ExpireIn = int64(0 - time.Since(auth.ExpireTime).Seconds())

	return
}

// Insert is ...
func (a *Authentication) Insert(auth Authentication, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "INSERT INTO session (courier_id, token, expire_time, create_time) VALUES (?,?,?,?)"
	query = db.Rebind(query)
	db.MustExec(query, auth.CourierID, auth.Token, auth.ExpireTime, auth.CreateTime)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Authentication][Insert] Error when commit to db : %s \n data : %v ", err, auth)
			err = errors.New("error when insert data to session")
		}
	}
	return err
}

// Update is ...
func (a *Authentication) Update(auth Authentication, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "UPDATE session SET courier_id = ?, token = ?, expire_time = ?, create_time = ? WHERE id = ?"
	query = db.Rebind(query)
	db.MustExec(query, auth.CourierID, auth.Token, auth.ExpireTime, auth.CreateTime, auth.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Authentication][Update] Error when commit to db : %s \n data : %v ", err, auth)
			err = errors.New("error when insert data to session")
		}
	}
	return err
}

// Remove is ...
func (a *Authentication) Remove(auth Authentication, tx *sqlx.Tx) (err error) {
	commitNow := false
	db := database.DB
	if tx == nil {
		tx = db.MustBegin()
		commitNow = true
	}

	query := "DELETE FROM session WHERE id = ?"
	query = db.Rebind(query)
	db.MustExec(query, auth.ID)

	if commitNow {
		err = tx.Commit()
		if err != nil {
			log.Printf("[Authentication][Update] Error when commit to db : %s \n data : %v ", err, auth)
			err = errors.New("error when insert data to session")
		}
	}
	return err
}
