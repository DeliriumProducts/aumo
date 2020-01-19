package aumo

import (
	"context"

	"upper.io/db.v3/lib/sqlbuilder"
)

// Tx represents an SQL transaction
// NOTE: this is an implementation detail (sql)
// How can this be solved?
type Tx interface {
	sqlbuilder.Database
}

// TxDo represents a helper function for making transactions
func TxDo(ctx context.Context, db sqlbuilder.Database, f func(sqlbuilder.Tx) error) error {
	tx, err := db.NewTx(ctx)
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

	return f(tx)
}
