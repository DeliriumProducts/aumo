package aumo

import "golang.org/x/crypto/bcrypt"

// User represents a user of aumo
type User struct {
	ID       uint      `json:"id" db:"id,omitempty"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"-" db:"password"`
	Avatar   string    `json:"avatar" db:"avatar"`
	Points   float64   `json:"points" db:"points"`
	Orders   []Order   `json:"orders" db:"-"`
	Receipts []Receipt `json:"receipts" db:"-"`
}

// ClaimReceipt claims a receipt and adds it to the receipts array
func (u *User) ClaimReceipt(r Receipt) {
	u.Receipts = append(u.Receipts, r)
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

// BuyItem adds the passed Product to the user's inventory
// if they have enough money to buy the desired quantity;
// substracts points from the user
func (u *User) BuyItem(p Product, quantity uint) error {
	total := p.Price * float64(quantity)

	// Check if the user has enough points
	if u.Points-total < 0 {
		return ErrNotSufficientPoints
	}

	// Check if there is enough in stock
	if p.Stock-quantity < 0 {
		return ErrNotInStock
	}

	// Substract the points of the user
	u.Points -= total

	o := NewOrder(u.ID, p.ID, &p)

	// Add the item to the orders array
	u.Orders = append(u.Orders, *o)
	return nil
}

// NewUser is a constructor for `User`
func NewUser(name string, email string, password string, avatar string) (*User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
		Avatar:   avatar,
	}, nil
}

// UserService contains all `User`
// related business logic
type UserService interface {
	User(id uint) (*User, error)
	Users() ([]User, error)
	Create(*User) error
	Update(id uint, u *User) error
	Delete(id uint) error
	ClaimReceipt(u *User, r *Receipt) error
}
