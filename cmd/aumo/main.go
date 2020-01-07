package main

import (
	"fmt"

	"github.com/deliriumproducts/aumo/mysql"
	upper "upper.io/db.v3/mysql"
)

func main() {
	PORT := 3000

	fmt.Printf("🧾 aumo server running on port %d\n", PORT)

	db, err := upper.Open(upper.ConnectionURL{
		User:     "root",
		Password: "fr3fou123/",
		Host:     "localhost",
		Database: "aumo",
	})

	defer db.Close()
	if err != nil {
		panic(err)
	}

	mysql.ExecSchema(db)
	mysql.NewProductService(db)
}
