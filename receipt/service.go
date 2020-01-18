package receipt

import (
	"errors"

	"github.com/deliriumproducts/aumo"
)

type receiptService struct {
	store     aumo.ReceiptStore
	userStore aumo.UserStore
}

// New returns an instance of `aumo.ReceiptService`
func New(store aumo.ReceiptStore) aumo.ReceiptService {
	return &receiptService{
		store: store,
	}
}

func (rs *receiptService) Receipt(id uint) (*aumo.Receipt, error) {
	return rs.store.FindByID(id)
}

func (rs *receiptService) Receipts() ([]aumo.Receipt, error) {
	return rs.store.FindAll()
}

func (rs *receiptService) Create(r *aumo.Receipt) error {
	return rs.store.Save(r)
}

func (rs *receiptService) Update(id uint, r *aumo.Receipt) error {
	return rs.store.Update(id, r)
}

func (rs *receiptService) Delete(id uint) error {
	return rs.store.Delete(id)
}

func (rs *receiptService) ClaimReceipt(uID uint, rID uint) (*aumo.Receipt, error) {
	receipt, err := rs.store.FindByID(rID)
	if err != nil {
		return nil, err
	}

	_, err = rs.userStore.FindByID(uID, false)
	if err != nil {
		if errors.Is(err, aumo.ErrReceiptUserNotExist) {
			return nil, aumo.ErrReceiptUserNotExist
		}

		return nil, err
	}

	err = receipt.Claim(uID)
	if err != nil {
		return nil, err
	}

	return receipt, rs.store.Update(rID, receipt)
}
