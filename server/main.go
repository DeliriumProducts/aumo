package main

import (
	"fmt"

	"github.com/fr3fou/aumo/server/aumo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	a := aumo.New(aumo.Config{})

	fmt.Printf("%+v", a)
	err = godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
}
