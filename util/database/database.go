package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	//postgres implementation
	_ "github.com/go-sql-driver/mysql"
)

var (
	//host
	host = "172.25.230.81"
	// username
	dbUser = "apps"
	// password
	dbPassword = "Apps*2013"
	// database
	dbName = "cakman_db"

	port = "3306"

	// DB Object
	DB *sqlx.DB
)

func InitDB() {

	dbinfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, host, port, dbName)
	log.Println(dbinfo)
	var err error
	DB, err = sqlx.Connect("mysql", dbinfo)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
