package province

import (
	"bytes"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rtawormy14/cakman-go/util/database"
	//postgre implementation
	_ "github.com/go-sql-driver/mysql"
)

// Provincer is Interface for province
type Provincer interface {
	GetProvince(int64) (Province, error)
	GetProvinceList(int64, int64, Province) ([]Province, error)
	Update(*sqlx.Tx) error
	Insert(*sqlx.Tx) error
	Remove(*sqlx.Tx) error
}

type Province struct {
	Code        int64  `db:"province_code" json:"province_code"`
	Name        string `db:"province_name" json:"province_name"`
	CountryCode int64  `db:"country_code" json:"-"`
}

func NewProvince() Provincer {
	return &Province{}
}

func (p *Province) GetProvince(code int64) (province Province, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return province, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT province_code, province_name FROM province WHERE province_code = ? LIMIT 1")

	query := db.Rebind(queryBuffer.String())
	err = db.Get(&province, query, code)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Province][GetProvince] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return
}

// GetProvinceList will return a list of provinces
func (p *Province) GetProvinceList(page, limit int64, filter Province) (provinces []Province, err error) {
	db := database.DB
	provinces = make([]Province, 0)

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT province_code, province_name FROM province ")
	queryBuffer.WriteString("WHERE country_code = ? AND province_name LIKE '%' || ? || '%' ")
	queryBuffer.WriteString("ORDER BY province_name ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET ? LIMIT ?")
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&provinces, query, filter.CountryCode, filter.Name, page, limit)
	} else {
		query := db.Rebind(queryBuffer.String())
		err = db.Select(&provinces, query, filter.CountryCode, filter.Name)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Printf("[Province][GetProvinceList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return provinces, nil
}

// Update is ...
func (p *Province) Update(tx *sqlx.Tx) error {
	return nil
}

// Insert is ...
func (p *Province) Insert(tx *sqlx.Tx) error {
	return nil
}

// Remove is ...
func (p *Province) Remove(tx *sqlx.Tx) error {
	return nil
}
