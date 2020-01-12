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

	Address := os.Getenv("ADDRESS")
	CookieSecret := os.Getenv("COOKIE_SECRET")
	RedisURL := os.Getenv("REDIS_URL")
	MySQLUser := os.Getenv("MYSQL_USER")
	MySQLPassword := os.Getenv("MYSQL_PASSWORD")
	MySQLHost := os.Getenv("MYSQL_HOST")
	MySQLDatabase := os.Getenv("MYSQL_DATABASE")

	db, err := upper.Open(upper.ConnectionURL{
		User:     MySQLUser,
		Password: MySQLPassword,
		Host:     MySQLHost,
		Database: MySQLDatabase,
	})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	conn, err := redis.DialURL(RedisURL)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = mysql.ExecSchema(db)
	if err != nil {
		panic(err)
	}

	ps := mysql.NewProductService(db)
	os := mysql.NewOrderService(db)
	rs := mysql.NewReceiptService(db)
	us := mysql.NewUserService(db, rs, ps, os)
	auth := auth.New(conn, us, 60*60*24)

	r := rest.New(rest.Config{
		UserService:    us,
		ReceiptService: rs,
		OrderService:   os,
		ProductService: ps,
		Auth:           auth,
		MountRoute:     "/api/v1",
		CookieSecret:   []byte(CookieSecret),
	})

	fmt.Printf("🧾 aumo server running on %s\n", Address)
	if err := http.ListenAndServe(Address, r); err != nil {
		panic(err)
	}
}
