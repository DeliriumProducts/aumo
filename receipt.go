package aumo

import "errors"

var (
	// ErrUserAlreadySet is an error for when a user has already claimed a receipt
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

// Receipt is a digital receipt
type Receipt struct {
	Content string `json:"content"`
	UserID  uint   `json:"-"`
}

// SetUser sets the user field of a receipt
func (r *Receipt) SetUser(userID uint) error {
	if r.UserID != 0 {
		return ErrUserAlreadySet
	}

	r.UserID = userID
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
