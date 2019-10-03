package main

import (
	"fmt"
	"os"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}

	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")

	MYSQL_STRING := MYSQL_USER + ":" + MYSQL_PASSWORD + "@(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE

	db, err := gorm.Open("mysql", MYSQL_STRING)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&aumo.User{})

	a := aumo.New(aumo.Config{})
	fmt.Printf("%+v", a)
}
