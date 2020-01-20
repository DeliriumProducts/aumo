package tests

import (
	"log"
	"os"

	"github.com/deliriumproducts/aumo/mysql"
	"github.com/joho/godotenv"
	"upper.io/db.v3/lib/sqlbuilder"
	upper "upper.io/db.v3/mysql"
)

// SetupDB creates a new in memory sqlite database
func SetupDB() (sqlbuilder.Database, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(".env file not found, reading directly from env variables")
	}

	db, err := upper.Open(upper.ConnectionURL{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Database: "aumo_test",
	})
	if err != nil {
		log.Fatalln("Couldn't establish a connection: ", err)
	}

	defer db.Close()

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
