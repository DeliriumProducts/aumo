package mysql

import (
	"github.com/deliriumproducts/aumo"
	"github.com/jinzhu/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	db.AutoMigrate(&aumo.Product{}, &aumo.User{}, &aumo.Order{}, &aumo.Receipt{})
}
