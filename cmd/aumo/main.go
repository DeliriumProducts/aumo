package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/net/http/rest"
	"github.com/deliriumproducts/aumo/product"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/deliriumproducts/aumo/user"
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

	ps := mysql.NewProductStore(db)
	_ = mysql.NewOrderStore(db)
	rs := mysql.NewReceiptStore(db)
	us := mysql.NewUserStore(db)
	auth := auth.New(conn, us, 60*60*24)

	r := rest.New(rest.Config{
		UserService:    user.New(us),
		ReceiptService: receipt.New(rs),
		OrderService:   nil,
		ProductService: product.New(ps),
		Auth:           auth,
		MountRoute:     "/api/v1",
	})

	fmt.Printf("🧾 aumo server running on %s\n", Address)
	if err := http.ListenAndServe(Address, r); err != nil {
		panic(err)
	}
}
