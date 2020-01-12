package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// OrderTable is the MySQL table for holding orders
const OrderTable = "orders"

type orderService struct {
	db sqlbuilder.Database
}

// NewOrderService returns a mysql instance of `aumo.OrderService`
func NewOrderService(db sqlbuilder.Database) aumo.OrderService {
	return &orderService{
		db: db,
	}
}

func (o *orderService) Order(id uint) (*aumo.Order, error) {
	os := &aumo.Order{}
	return os, o.db.Collection(OrderTable).Find("id", id).One(os)
}

func (o *orderService) Orders() ([]aumo.Order, error) {
	oss := []aumo.Order{}
	return oss, o.db.Collection(OrderTable).Find().All(&oss)
}

func (o *orderService) Create(os *aumo.Order) error {
	return o.db.Collection(OrderTable).InsertReturning(os)
}

func (o *orderService) Update(id uint, or *aumo.Order) error {
	return o.db.Collection(OrderTable).Find("id", id).Update(or)
}

func (o *orderService) Delete(id uint) error {
	return o.db.Collection(OrderTable).Find("id", id).Delete()
}
