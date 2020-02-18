package aumo

import (
	"github.com/google/uuid"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Receipt is a digital receipt
type Receipt struct {
	ReceiptID uuid.UUID  `json:"receipt_id" db:"receipt_id"`
	Content   string     `json:"content" db:"content"`
	UserID    *uuid.UUID `json:"-" db:"user_id,omitempty"`
	ShopID    uint       `json:"shop_id" db:"shop_id"`
	Shop      *Shop      `json:"shop" db:"-"`
}

// NewReceipt is a contrsuctor for `Receipt`
func NewReceipt(content string, sID uint) *Receipt {
	return &Receipt{
		ReceiptID: uuid.New(),
		Content:   content,
		ShopID:    sID,
	}
}

// Claim sets the user field of a receipt
func (r *Receipt) Claim(uID uuid.UUID) error {
	if r.IsClaimed() {
		return ErrUserAlreadySet
	}

	r.UserID = &uID
	return nil
}

// IsClaimed checks if the Receipt has been claimed
func (r *Receipt) IsClaimed() bool {
	return r.UserID != nil
}

// ReceiptService contains all `Receipt`
// related business logic
type ReceiptService interface {
	Receipt(id string) (*Receipt, error)
	Receipts() ([]Receipt, error)
	Create(*Receipt) error
	Update(id string, r *Receipt) error
	Delete(id string) error
	ClaimReceipt(uID, rID string) (*Receipt, error)
}

// ReceiptStore contains all `Receipt`
// related persistence logic
type ReceiptStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id string) (*Receipt, error)
	FindAll(tx Tx) ([]Receipt, error)
	Save(tx Tx, r *Receipt) error
	Update(tx Tx, id string, r *Receipt) error
	Delete(tx Tx, id string) error
}
