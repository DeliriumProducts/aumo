package aumo

import "errors"

var (
	// ErrUserAlreadySet is an error for when a user has already claimed a receipt
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

// Receipt is a digital receipt
type Receipt struct {
	ID      uint   `json:"id" db:"id"`
	Content string `json:"content" db:"content"`
	UserID  uint   `json:"-" db:"user_id"`
}

// NewReceipt is a contrsuctor for `Receipt`
func NewReceipt(uid uint, content string) *Receipt {
	return &Receipt{
		Content: content,
		UserID:  uid,
	}
}

// SetUser sets the user field of a receipt
func (r *Receipt) SetUser(uid uint) error {
	if r.UserID != 0 {
		return ErrUserAlreadySet
	}

	r.UserID = uid
	return nil
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
