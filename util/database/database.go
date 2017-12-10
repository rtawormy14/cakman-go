package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	//postgres implementation
	_ "github.com/lib/pq"
)

var (
	//host
	host = "localhost"
	// username
	dbUser = "cakman"
	// password
	dbPassword = "password"
	// database
	dbName = "cakman_db"

	// DB Object
	DB *sqlx.DB
)

func InitDB() {
	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		host, dbUser, dbPassword, dbName)
	var err error
	DB, err = sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
