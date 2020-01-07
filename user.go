package aumo

import "golang.org/x/crypto/bcrypt"

// User represents a user of aumo
type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
	Avatar   string    `json:"avatar"`
	Points   float64   `json:"points"`
	Orders   []Product `json:"orders"`
	Receipts []Receipt `json:"receipts"`
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
}
