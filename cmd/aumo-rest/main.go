package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/mail"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/net/http/rest"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/deliriumproducts/aumo/products"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/deliriumproducts/aumo/shops"
	"github.com/deliriumproducts/aumo/users"
	"github.com/deliriumproducts/aumo/verifications"
	"github.com/go-redis/redis/v7"
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
	RedisDatabase := os.Getenv("REDIS_DATABASE")
	MySQLUser := os.Getenv("MYSQL_USER")
	MySQLPassword := os.Getenv("MYSQL_PASSWORD")
	MySQLHost := os.Getenv("MYSQL_HOST")
	MySQLDatabase := os.Getenv("MYSQL_DATABASE")
	InitialAdminPassword := os.Getenv("INITIAL_ADMIN_PASSWORD")
	FrontendURL := os.Getenv("FRONTEND_URL")
	Env := os.Getenv("ENV")
	BackendURL := os.Getenv("BACKEND_URL")
	SMTPHost := os.Getenv("SMTP_HOST")
	SMTPUser := os.Getenv("SMTP_USER")
	SMTPPass := os.Getenv("SMTP_PASS")

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

	redisDbN, err := strconv.Atoi(RedisDatabase)
	if err != nil {
		panic(err)
	}

	conn := redis.NewClient(&redis.Options{
		Addr: RedisURL,
		DB:   redisDbN,
	})

	err = conn.Ping().Err()
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	err = mysql.ExecSchema(db)
	if err != nil {
		panic(err)
	}

	var mailer mail.Mailer
	if Env == "PROD" {
		mailer = &mail.ProdMailer{
			Host: SMTPHost,
			Port: "587",
			User: SMTPUser,
			Pass: SMTPPass,
			From: "thedeliriumproducts@gmail.com",
		}
	} else {
		mailer = &mail.DevMailer{}
	}

	ps := mysql.NewProductStore(db)
	os := mysql.NewOrderStore(db)
	rs := mysql.NewReceiptStore(db)
	us := mysql.NewUserStore(db)
	ss := mysql.NewShopStore(db)
	so := mysql.NewShopOwnersStore(db)

	auth := auth.New(conn, us, FrontendURL, "/", time.Hour*24)

	_, err = users.InitialAdmin(us, InitialAdminPassword, "admin@deliriumproducts.me")
	if err != nil {
		panic(err)
	}

	r := rest.New(&rest.Config{
		UserService:    users.New(us, so),
		ReceiptService: receipt.New(rs, us),
		OrderService:   ordering.New(os, ps, us),
		ProductService: products.New(ps),
		ShopService:    shops.New(ss, so, us),
		Auth:           auth,
		MountRoute:     "/api/v1",
		BackendURL:     BackendURL,
		Verifier:       verifications.New(mailer, conn),
	})

	fmt.Printf("🧾 aumo server running on %s\n", Address)
	if err := http.ListenAndServe(Address, r); err != nil {
		panic(err)
	}
}
