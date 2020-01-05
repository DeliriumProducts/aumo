package mysql

import (
	"github.com/deliriumproducts/aumo"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

// New returns a Context instance while also migrating all the structs from luncherbox
func New(dialect string, args ...interface{}) (*DB, error) {
	db, err := gorm.Open(dialect, args...)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&aumo.Product{}, &aumo.User{}, &aumo.Receipt{})

	return &DB{
		db,
	}, nil
}
