package mysql

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	upper "upper.io/db.v3"
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

func (r *receiptStore) FindByID(tx aumo.Tx, id string) (*aumo.Receipt, error) {
	var err error
	receipt := &aumo.Receipt{}

	if tx == nil {
		tx, err = r.db.NewTx(context.Background())

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	err = tx.Collection(ReceiptTable).Find("receipt_id", id).One(receipt)

	switch {
	case err == nil:
		break
	case errors.Is(err, upper.ErrNoMoreRows):
		return nil, aumo.ErrReceiptNotFound
	default:
		return nil, err
	}

	return receipt, err
}

func (r *receiptStore) FindAll(tx aumo.Tx) ([]aumo.Receipt, error) {
	var err error
	receipts := []aumo.Receipt{}

	if tx == nil {
		tx, err = r.db.NewTx(context.Background())

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return receipts, tx.Collection(ReceiptTable).Find().All(&receipts)
}

func (r *receiptStore) Save(tx aumo.Tx, rs *aumo.Receipt) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return tx.Collection(ReceiptTable).InsertReturning(rs)
}

func (r *receiptStore) Update(tx aumo.Tx, id string, rr *aumo.Receipt) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return tx.Collection(ReceiptTable).Find("receipt_id", id).Update(rr)
}

func (r *receiptStore) Delete(tx aumo.Tx, id string) error {
	var err error

	if tx == nil {
		tx, err = r.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return r.db.Collection(ReceiptTable).Find("receipt_id", id).Delete()
}
