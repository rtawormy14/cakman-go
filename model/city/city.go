package city

import (
	"bytes"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rtawormy14/cakman-go/util/database"
	//postgre implementation
	_ "github.com/lib/pq"
)

// Citier interface
type Citier interface {
	GetCity(int64) (City, error)
	GetCityList(int64, int64, City) ([]City, error)
	Update(*sqlx.Tx) error
	Insert(*sqlx.Tx) error
	Remove(*sqlx.Tx) error
}

// City struct
type City struct {
	Code         int64  `db:"city_code" json:"city_code"`
	Name         string `db:"city_name" json:"city_name"`
	ProvinceCode int64  `db:"province_code" json:"-"`
}

// New City
func NewCity() Citier {
	return &City{}
}

// GetCity will return city object if exist
func (p *City) GetCity(code int64) (city City, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return city, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT city_code, city_name FROM city WHERE city_code = $1 LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&city, query, code)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[City][GetCity] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetCityList will return a list of city 	based on parameter defined
func (p *City) GetCityList(page, limit int64, filter City) (cities []City, err error) {
	db := database.DB

	cities = make([]City, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT city_code, city_name FROM city ")
	queryBuffer.WriteString("WHERE province_code = $1 AND city_name ILIKE '%' || $2 || '%' ")
	queryBuffer.WriteString("ORDER BY city_name ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET $3 LIMIT $4")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&cities, query, filter.ProvinceCode, filter.Name, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&cities, query, filter.ProvinceCode, filter.Name)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[City][GetCityList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// Update will update city data
func (p *City) Update(tx *sqlx.Tx) error {
	return nil
}

// Insert will insert city data
func (p *City) Insert(tx *sqlx.Tx) error {
	return nil
}

// Remove will remove city data
func (p *City) Remove(tx *sqlx.Tx) error {
	return nil
}
