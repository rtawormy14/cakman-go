package country

import (
	"bytes"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	//postgre implementation
	_ "github.com/lib/pq"

	"github.com/rtawormy14/cakman-go/util/database"
)

// Countrier is interface for Country
type Countrier interface {
	GetCountry(int64)
	GetCountryList(int64, int64, Country)
	Update(Country, *sqlx.Stmt)
	Insert(Country, *sqlx.Stmt)
	Remove(Country, *sqlx.Stmt)
}

// Country is data struct for Country
type Country struct {
	Code int64  `db:"country_code" json:"country_code"`
	Name string `db:"country_name" json:"country_Name"`
}

// New Country
func New() Country {
	return Country{}
}

// Get Country Object
func GetCountry(code int64) (country Country, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return country, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT country_code, country_name FROM country WHERE country_code = $1 LIMIT 1")

	err = db.Get(&country, queryBuffer.String(), code)
	if err != nil {
		log.Printf("[Country][GetCountry] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetCountryList will return list of country.
// Default page = 0. Default limit = no-limit.
func (p *Country) GetCountryList(page, limit int64, filter Country) (countries []Country, err error) {
	db := database.DB
	countries = make([]Country, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT country_code, country_name FROM country WHERE country_name ILIKE '%' || $1 || '%' ORDER BY country_name ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET $2 LIMIT $3")
		err = db.Select(&countries, queryBuffer.String(), filter.Name, page, limit)
	} else {
		err = db.Select(&countries, queryBuffer.String(), filter.Name)
	}

	if err != nil {
		log.Printf("[Country][GetCountryList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return

}

func (p *Country) Update(country Country, tx *sqlx.Stmt) error {
	return nil
}

func (p *Country) Insert(country Country, tx *sqlx.Stmt) error {
	return nil
}

func (p *Country) Remove(country Country, tx *sqlx.Stmt) error {
	return nil
}
