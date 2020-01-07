package tests

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	upper "upper.io/db.v3/mysql"
)

// SetupDB creates a new in memory sqlite database
func SetupDB() (sqlbuilder.Database, error) {
	return upper.Open(upper.ConnectionURL{
		User:     "root",
		Password: "fr3fou123/",
		Host:     "localhost",
		Database: "aumo_test",
	})
}

// TruncateTables deletes all rows from the provided tables, while manintaing the schema
func TruncateTables(db sqlbuilder.Database, tables ...string) {
	for _, table := range tables {
		_, err := db.Exec(`TRUNCATE TABLE ?`, table)
		log.Println(err)
	}
}
