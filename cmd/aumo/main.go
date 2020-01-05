package main

import (
	"fmt"

	"github.com/deliriumproducts/aumo/db"
	"github.com/deliriumproducts/aumo/mysql"
)

func main() {
	PORT := 3000
	fmt.Printf("🧾 aumo server running on port %d\n", PORT)
	MYSQL_STRING := "root" + ":" + "fr3fou123/" + "@(" + "localhost" + ":" + "3306" + ")/" + "aumo" + "?parseTime=true"
	_, err := mysql.New("mysql", MYSQL_STRING)
	if err != nil {
		panic(err)
	}
	db.Users()
}
