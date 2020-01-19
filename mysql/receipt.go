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

func (r *receiptStore) FindByID(tx aumo.Tx, id uint) (*aumo.Receipt, error) {
	if tx == nil {
		tx = r.db
	}
	receipt := &aumo.Receipt{}
	return receipt, tx.Collection(ReceiptTable).Find("receipt_id", id).One(receipt)
}

func (r *receiptStore) FindAll(tx aumo.Tx) ([]aumo.Receipt, error) {
	if tx == nil {
		tx = r.db
	}
	receipts := []aumo.Receipt{}
	return receipts, tx.Collection(ReceiptTable).Find().All(&receipts)
}

func (r *receiptStore) Save(tx aumo.Tx, rs *aumo.Receipt) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Collection(ReceiptTable).InsertReturning(rs)
}

func (r *receiptStore) Update(tx aumo.Tx, id uint, rr *aumo.Receipt) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Collection(ReceiptTable).Find("receipt_id", id).Update(rr)
}

func (r *receiptStore) Delete(tx aumo.Tx, id uint) error {
	if tx == nil {
		tx = r.db
	}
	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Delete()
}
