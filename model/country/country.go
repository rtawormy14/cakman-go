package country

import (
	"bytes"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/go-sql-driver/mysql"

	"github.com/rtawormy14/cakman-go/util/database"
)

// Countrier is interface for Country
type Countrier interface {
	GetCountry(int64) (Country, error)
	GetCountryList(int64, int64, Country) ([]Country, error)
	Update(*sqlx.Tx) error
	Insert(*sqlx.Tx) error
	Remove(*sqlx.Tx) error
}

// Country is data struct for Country
type Country struct {
	Code int64  `db:"country_code" json:"country_code"`
	Name string `db:"country_name" json:"country_Name"`
}

// New Country
func NewCountry() Countrier {
	return &Country{}
}

// Get Country Object
func (c *Country) GetCountry(code int64) (country Country, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return country, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT country_code, country_name FROM country WHERE country_code = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&country, query, code)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Country][GetCountry] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetCountryList will return list of country.
// Default page = 0. Default limit = no-limit.
func (c *Country) GetCountryList(page, limit int64, filter Country) (countries []Country, err error) {
	db := database.DB
	countries = make([]Country, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT country_code, country_name FROM country WHERE country_name LIKE '%' || ? || '%' ORDER BY country_name ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET ? LIMIT ?")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&countries, query, filter.Name, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&countries, query, filter.Name)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Country][GetCountryList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return

}

// Update is ...
func (c *Country) Update(tx *sqlx.Tx) error {
	return nil
}

// Insert is ...
func (c *Country) Insert(tx *sqlx.Tx) error {
	return nil
}

// Remove is ...
func (c *Country) Remove(tx *sqlx.Tx) error {
	return nil
}
