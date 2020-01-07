package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deliriumproducts/aumo/mysql"
	"github.com/joho/godotenv"
	upper "upper.io/db.v3/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, reading directly from env variables")
	}

	ADDRESS := os.Getenv("ADDRESS")

	fmt.Printf("🧾 aumo server running on %s\n", ADDRESS)

	db, err := upper.Open(upper.ConnectionURL{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Database: os.Getenv("MYSQL_DATABASE"),
	})

	defer db.Close()
	if err != nil {
		panic(err)
	}

	mysql.ExecSchema(db)
}
