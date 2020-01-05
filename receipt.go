package aumo

import "errors"

var (
	ErrUserAlreadySet = errors.New("aumo: this receipt has already been claimed")
)

// Receipt is a digital receipt
type Receipt struct {
	Model   `json:"model"`
	Content string `json:"content"`
	User    *User  `json:"-"`
}

// SetUser sets the user field of a receipt
func (r *Receipt) SetUser(u *User) error {
	if r.User != nil {
		return ErrUserAlreadySet
	}

	r.User = u
	return nil
}

// ReceiptService contains all `Receipt`
// related business logic
type ReceiptService interface {
	Receipt(id uint) (*Receipt, error)
	Receipts() ([]Receipt, error)
	Save(*Receipt) error
	Update(*Receipt) error
	Delete(*Receipt) error
}
