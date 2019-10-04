package main

import (
	"fmt"
	"os"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/fr3fou/aumo/server/web"
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
	ADDRESS := os.Getenv("ADDRESS")

	MYSQL_STRING := MYSQL_USER + ":" + MYSQL_PASSWORD + "@(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?parseTime=true"

	db, err := gorm.Open("mysql", MYSQL_STRING)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	a := aumo.New(aumo.Config{
		DB: db,
	})

	w := web.New(web.Config{
		Aumo: a,
	})

	log.Println("Aumo server running on port ", PORT)
	if err := http.ListenAndServe("localhost:"+PORT, w.Router); err != nil {
		panic(err)
	}
}
