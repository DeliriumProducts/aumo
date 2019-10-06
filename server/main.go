package main

import (
	"log"
	"net/http"
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
		log.Println(".env file not found, reading directly from env variables")
	}

	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_HOST := os.Getenv("MYSQL_HOST")
	MYSQL_PORT := os.Getenv("MYSQL_PORT")
	ADDRESS := os.Getenv("ADDRESS")
	COOKIE_SECRET := os.Getenv("COOKIE_SECRET")

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
		Aumo:         a,
		CookieSecret: []byte(COOKIE_SECRET),
	})

	log.Println("Aumo server running on port", ADDRESS)
	if err := http.ListenAndServe(ADDRESS, w.Router); err != nil {
		panic(err)
	}
}
