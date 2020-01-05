package main

// go:generate sqlboiler mysql -p "dbx" -o "dbx" --wipe

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	PORT := 3000
	fmt.Printf("🧾 aumo server running on port %d\n", PORT)
	MYSQL_STRING := "root" + ":" + "fr3fou123/" + "@(" + "localhost" + ":" + "3306" + ")/" + "aumo" + "?parseTime=true"
	db, err := sql.Open("mysql", MYSQL_STRING)
	if err != nil {
		panic(err)
	}

}
