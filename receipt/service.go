package receipt

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	store     aumo.ReceiptStore
	userStore aumo.UserStore
}

// New returns an instance of `aumo.ReceiptService`
func New(store aumo.ReceiptStore, userStore aumo.UserStore) aumo.ReceiptService {
	return &service{
		store:     store,
		userStore: userStore,
	}
}

func (rs *service) Receipt(id string) (*aumo.Receipt, error) {
	return rs.store.FindByID(nil, id)
}

func (rs *service) Receipts() ([]aumo.Receipt, error) {
	return rs.store.FindAll(nil)
}

func (rs *service) Create(r *aumo.Receipt) error {
	return rs.store.Save(nil, r)
}

func (rs *service) Update(id string, r *aumo.Receipt) error {
	return rs.store.Update(nil, id, r)
}

func (rs *service) Delete(id string) error {
	return rs.store.Delete(nil, id)
}

func (rs *service) ClaimReceipt(uID string, rID string) (*aumo.Receipt, error) {
	var receipt *aumo.Receipt
	db := rs.store.DB()

	err := aumo.TxDo(context.Background(), db, func(tx sqlbuilder.Tx) error {
		var err error
		receipt, err = rs.store.FindByID(tx, rID)
		if err != nil {
			return err
		}

		user, err := rs.userStore.FindByID(tx, uID, false)
		if err != nil {
			return err
		}

		uuid, err := uuid.Parse(uID)
		if err != nil {
			return err
		}

		err = receipt.Claim(uuid)
		if err != nil {
			return err
		}

		err = rs.store.Update(tx, rID, receipt)
		if err != nil {
			return err
		}

		user.Points += aumo.UserPointsPerReceipt

		err = rs.userStore.Update(tx, uID, user)
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
