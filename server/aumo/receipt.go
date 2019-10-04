package aumo

import (
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrUserIDAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

type Receipt struct {
	gorm.Model
	Content string
	UserID  sql.NullInt64
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

	if err := a.DB.Create(receipt).Error; err != nil {
		return Receipt{}, err
	}
	return *receipt, nil
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
