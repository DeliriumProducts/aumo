package main

// go:generate sqlboiler mysql -p "dbx" -o "dbx" --wipe

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/deliriumproducts/aumo/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/boil"
)

func main() {
	PORT := 3000
	fmt.Printf("🧾 aumo server running on port %d\n", PORT)
	MYSQL_STRING := "root" + ":" + "fr3fou123/" + "@(" + "localhost" + ":" + "3306" + ")/" + "aumo" + "?parseTime=true"
	d, err := sql.Open("mysql", MYSQL_STRING)
	if err != nil {
		panic(err)
	}

	user := db.User{Name: "pesho"}
	user.Insert(context.Background(), d, boil.Infer())
	fmt.Println(user.ID)

	user.AddOrders(context.Background(), d, true, &db.Order{UserID: user.ID, ProductID: 2})
}
