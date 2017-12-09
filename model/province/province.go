package province

import (
	"bytes"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/rtawormy14/cakman-go/util/database"
	//postgre implementation
	_ "github.com/lib/pq"
)

// Provincer is Interface for province
type Provincer interface {
	GetProvince(int64)
	GetProvinceList(int64, int64, Province)
	Update(Province, *sqlx.Stmt)
	Insert(Province, *sqlx.Stmt)
	Remove(Province, *sqlx.Stmt)
}

type Province struct {
	Code        int64  `db:"province_code" json:"province_code"`
	Name        string `db:"province_name" json:"province_name"`
	CountryCode int64  `db:"country_code" json:"-"`
}

func New() Province {
	return Province{}
}

func GetProvince(code int64) (province Province, err error) {
	db := database.DB

	//validate parameter
	if code <= 0 {
		return province, errors.New("code is not valid")
	}

	var queryBuffer bytes.Buffer
	queryBuffer.WriteString("SELECT province_code, province_name FROM province WHERE province_code = $1 LIMIT 1")

	err = db.Get(&province, queryBuffer.String(), code)
	if err != nil {
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
	queryBuffer.WriteString("WHERE country_code = $1 AND province_name ILIKE '%' || $2 || '%' ")
	queryBuffer.WriteString("ORDER BY province_name ASC ")

	if limit > 0 {
		queryBuffer.WriteString("OFFSET $3 LIMIT $4")
		err = db.Select(&provinces, queryBuffer.String(), filter.CountryCode, filter.Name, page, limit)
	} else {
		err = db.Select(&provinces, queryBuffer.String(), filter.CountryCode, filter.Name)
	}

	if err != nil {
		log.Printf("[Province][GetProvinceList] error while querying : \n %v \n %v", queryBuffer.String(), err)
	}
	return provinces, nil
}

func (p *Province) Update(province Province, tx *sqlx.Stmt) error {
	return nil
}

func (p *Province) Insert(province Province, tx *sqlx.Stmt) error {
	return nil
}

func (p *Province) Remove(province Province, tx *sqlx.Stmt) error {
	return nil
}
