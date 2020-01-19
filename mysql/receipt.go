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

func (r *receiptStore) DB() sqlbuilder.Database {
	return r.db
}

func (r *receiptStore) FindByID(tx aumo.Tx, id uint) (*aumo.Receipt, error) {
	var err error
	receipt := &aumo.Receipt{}

	if tx == nil {
		tx, err = r.db.NewTx(nil)

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return receipt, tx.Collection(ReceiptTable).Find("receipt_id", id).One(receipt)
}

func (r *receiptStore) FindAll(tx aumo.Tx) ([]aumo.Receipt, error) {
	var err error
	receipts := []aumo.Receipt{}

	if tx == nil {
		tx, err = r.db.NewTx(nil)

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return receipts, tx.Collection(ReceiptTable).Find().All(&receipts)
}

func (r *receiptStore) Save(tx aumo.Tx, rs *aumo.Receipt) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return tx.Collection(ReceiptTable).InsertReturning(rs)
}

func (r *receiptStore) Update(tx aumo.Tx, id uint, rr *aumo.Receipt) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return tx.Collection(ReceiptTable).Find("receipt_id", id).Update(rr)
}

func (r *receiptStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Delete()
}
