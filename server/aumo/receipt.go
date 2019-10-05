package aumo

import (
	"database/sql"
	"errors"
)

var (
	ErrUserIDAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

type Receipt struct {
	Model   `json:"-"`
	Content string        `json:"content"`
	UserID  sql.NullInt64 `json:"user_id"`
}

// SetUserID claims a receipt with the provided ID
func (r *Receipt) SetUserID(userID int64) error {
	if !r.UserID.Valid {
		return ErrUserIDAlreadySet
	}

	r.UserID.Int64 = userID

	return nil
}

// CreateReceipt creates a receipt
func (a *Aumo) CreateReceipt(content string) (Receipt, error) {
	receipt := &Receipt{
		Content: content,
	}

	if err := a.db.Create(receipt).Error; err != nil {
		return Receipt{}, err
	}
	return *receipt, nil
}

// GetReceiptByID gets a receipt by ID
func (a *Aumo) GetReceiptByID(id uint) (Receipt, error) {
	var r Receipt
	err := a.firstX(&r, "id = ?", id)
	return r, err
}

// SetReceiptUserID claims the receipt by calling the ClaimReceipt(r) (adds receipt to the receipt list of the user)
// Sets the user id in the receipt (receipt is claimed by the user)
func (a *Aumo) SetReceiptUserID(u User, r Receipt) error {
	u.ClaimReceipt(r)
	r.SetUserID(int64(u.ID))

	err := a.updateX(u)
	if err != nil {
		return err
	}

	err = a.updateX(r)
	if err != nil {
		return err
	}

	return nil
}
