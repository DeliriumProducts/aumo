package receipt

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	store aumo.ReceiptStore
}

// New returns an instance of `aumo.ReceiptService`
func New(store aumo.ReceiptStore) aumo.ReceiptService {
	return &service{
		store: store,
	}
}

func (rs *service) Receipt(id uint) (*aumo.Receipt, error) {
	return rs.store.FindByID(nil, id)
}

func (rs *service) Receipts() ([]aumo.Receipt, error) {
	return rs.store.FindAll(nil)
}

func (rs *service) Create(r *aumo.Receipt) error {
	return rs.store.Save(nil, r)
}

func (rs *service) Update(id uint, r *aumo.Receipt) error {
	return rs.store.Update(nil, id, r)
}

func (rs *service) Delete(id uint) error {
	return rs.store.Delete(nil, id)
}

func (rs *service) ClaimReceipt(uID uint, rID uint) (*aumo.Receipt, error) {
	var receipt *aumo.Receipt
	db := rs.store.DB()

	err := aumo.TxDo(context.Background(), db, func(tx sqlbuilder.Tx) error {
		var err error
		receipt, err = rs.store.FindByID(tx, rID)
		if err != nil {
			return err
		}

		err = receipt.Claim(uID)
		if err != nil {
			return err
		}

		err = rs.store.Update(tx, rID, receipt)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return receipt, err
}
