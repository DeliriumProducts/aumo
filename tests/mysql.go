package tests

import (
	"log"

	"github.com/deliriumproducts/aumo/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
	upper "upper.io/db.v3/mysql"
)

// SetupDB creates a new in memory sqlite database
func SetupDB() (sqlbuilder.Database, error) {
	db, err := upper.Open(upper.ConnectionURL{
		User:     "root",
		Password: "fr3fou123/",
		Host:     "localhost",
		Database: "aumo_test",
	})
	if err != nil {
		log.Fatalln("Couldn't establish a connection: ", err)
	}

	TidyDB(db)

	return db, err
}

// TidyDB drops the database and creates it again
func TidyDB(db sqlbuilder.Database) {
	_, err := db.Exec("DROP DATABASE IF EXISTS aumo_test")
	if err != nil {
		log.Fatalln("Couldn't DROP database: ", err)
	}

	_, err = db.Exec("CREATE DATABASE aumo_test")
	if err != nil {
		log.Fatalln("Couldn't CREATE database: ", err)
	}

	_, err = db.Exec("USE aumo_test")
	if err != nil {
		log.Fatalln("Couldn't USE database: ", err)
	}

	err = mysql.ExecSchema(db)
	if err != nil {
		log.Fatalln("Couldn't exec schema: ", err)
	}
}
