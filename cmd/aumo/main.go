package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/net/http/rest"
	"github.com/deliriumproducts/aumo/net/http/rest/auth"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	upper "upper.io/db.v3/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, reading directly from env variables")
	}

	ADDRESS := os.Getenv("ADDRESS")
	COOKIE_SECRET := os.Getenv("COOKIE_SECRET")

	db, err := upper.Open(upper.ConnectionURL{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Database: os.Getenv("MYSQL_DATABASE"),
	})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = mysql.ExecSchema(db)
	if err != nil {
		panic(err)
	}

	ps := mysql.NewProductService(db)
	os := mysql.NewOrderService(db)
	rs := mysql.NewReceiptService(db)
	us := mysql.NewUserService(db, rs, ps, os)

	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}

	auth := auth.New(auth.Config{
		Redis:      conn,
		ExpiryTime: "86400",
	})

	r := rest.New(rest.Config{
		UserService:    us,
		ReceiptService: rs,
		OrderService:   os,
		ProductService: ps,
		Auth:           auth,
		CookieSecret:   []byte(COOKIE_SECRET),
	})

	fmt.Printf("🧾 aumo server running on %s\n", ADDRESS)
	if err := http.ListenAndServe(ADDRESS, r.Router); err != nil {
		panic(err)
	}
}
