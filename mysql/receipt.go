package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ReceiptTable is the MySQL table for holding receipts
const ReceiptTable = "receipts"

type receiptStore struct {
	db sqlbuilder.Database
}

// NewReceiptStore returns a mysql instance of `aumo.ReceiptStore`
func NewReceiptStore(db sqlbuilder.Database) aumo.ReceiptStore {
	return &receiptStore{
		db: db,
	}
}

func (r *receiptStore) FindByID(id uint) (*aumo.Receipt, error) {
	receipt := &aumo.Receipt{}
	return receipt, r.db.Collection(ReceiptTable).Find("receipt_id", id).One(receipt)
}

func (r *receiptStore) FindAll() ([]aumo.Receipt, error) {
	receipts := []aumo.Receipt{}
	return receipts, r.db.Collection(ReceiptTable).Find().All(&receipts)
}

func (r *receiptStore) Save(rs *aumo.Receipt) error {
	return r.db.Collection(ReceiptTable).InsertReturning(rs)
}

func (r *receiptStore) Update(id uint, rr *aumo.Receipt) error {
	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Update(rr)
}

func (r *receiptStore) Delete(id uint) error {
	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Delete()
}
