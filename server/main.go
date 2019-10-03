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

	MYSQL_STRING := MYSQL_USER + ":" + MYSQL_PASSWORD + "@(" + MYSQL_HOST + ":" + MYSQL_PORT + ")/" + MYSQL_DATABASE + "?parseTime=true"

	db, err := gorm.Open("mysql", MYSQL_STRING)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	a := aumo.New(aumo.Config{
		DB: db,
	})

	a.CreateUser("fr3fou", "simo3003@me.com", "fr3fou123/")
	u, _ := a.GetUserByEmail("simo3003@me.com")
	si, _ := a.CreateShopItem("Pesho", 34, "pesho", 10)
	u.SetUserPoints(500000)
	fmt.Println(u.Points)

	fmt.Println(u.BuyItem(si, 10))
	fmt.Printf("%+v", u)
}
