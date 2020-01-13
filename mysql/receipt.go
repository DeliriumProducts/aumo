package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ReceiptTable is the MySQL table for holding receipts
const ReceiptTable = "receipts"

type receiptService struct {
	db sqlbuilder.Database
}

// NewReceiptService returns a mysql instance of `aumo.ReceiptService`
func NewReceiptService(db sqlbuilder.Database) aumo.ReceiptService {
	return &receiptService{
		db: db,
	}
}

func (r *receiptService) Receipt(id uint) (*aumo.Receipt, error) {
	rs := &aumo.Receipt{}
	return rs, r.db.Collection(ReceiptTable).Find("receipt_id", id).One(rs)
}

func (r *receiptService) Receipts() ([]aumo.Receipt, error) {
	rss := []aumo.Receipt{}
	return rss, r.db.Collection(ReceiptTable).Find().All(&rss)
}

func (r *receiptService) Create(rs *aumo.Receipt) error {
	return r.db.Collection(ReceiptTable).InsertReturning(rs)
}

func (r *receiptService) Update(id uint, rr *aumo.Receipt) error {
	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Update(rr)
}

func (r *receiptService) Delete(id uint) error {
	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Delete()
}

func (r *receiptService) ClaimReceipt(uID uint, rID uint) (*aumo.Receipt, error) {
	receipt, err := r.Receipt(rID)
	if err != nil {
		return nil, err
	}

	err = receipt.Claim(uID)
	if err != nil {
		return nil, err
	}

	return receipt, r.Update(rID, receipt)
}
