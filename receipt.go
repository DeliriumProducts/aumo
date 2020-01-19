package aumo

import (
	"errors"

	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	// ErrUserAlreadySet is an error for when a user has already claimed a receipt
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
	// ErrReceiptUserNotExist is an error for when a user doesn't exist when trying to claim a receipt
	ErrReceiptUserNotExist = errors.New("aumo: can't claim a receipt for a user that doesn't exist")
)

// Receipt is a digital receipt
type Receipt struct {
	ReceiptID uint   `json:"receipt_id" db:"receipt_id"`
	Content   string `json:"content" db:"content" validate:"required"`
	UserID    *uint  `json:"-" db:"user_id,omitempty"`
}

// NewReceipt is a contrsuctor for `Receipt`
func NewReceipt(content string) *Receipt {
	return &Receipt{
		Content: content,
	}
}

// Claim sets the user field of a receipt
func (r *Receipt) Claim(uid uint) error {
	if r.IsClaimed() {
		return ErrUserAlreadySet
	}

	r.UserID = &uid
	return nil
}

// IsClaimed checks if the Receipt has been claimed
func (r *Receipt) IsClaimed() bool {
	return r.UserID != nil
}

// ReceiptService contains all `Receipt`
// related business logic
type ReceiptService interface {
	Receipt(id uint) (*Receipt, error)
	Receipts() ([]Receipt, error)
	Create(*Receipt) error
	Update(id uint, r *Receipt) error
	Delete(id uint) error
	ClaimReceipt(uID uint, rID uint) (*Receipt, error)
}

// ReceiptStore contains all `Receipt`
// related persistence logic
type ReceiptStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id uint) (*Receipt, error)
	FindAll(tx Tx) ([]Receipt, error)
	Save(tx Tx, r *Receipt) error
	Update(tx Tx, id uint, r *Receipt) error
	Delete(tx Tx, id uint) error
}
