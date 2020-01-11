package aumo

import "errors"

var (
	// ErrUserAlreadySet is an error for when a user has already claimed a receipt
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

// Receipt is a digital receipt
type Receipt struct {
	ReceiptID uint   `json:"receiptID" db:"receipt_id"`
	Content   string `json:"content" db:"content"`
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
}
