package aumo

import "golang.org/x/crypto/bcrypt"

// User represents a user of aumo
type User struct {
	Model
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"unique;not null"`
	Password string    `json:"-" gorm:"not null"`
	Avatar   string    `json:"avatar" gorm:"not null"`
	Points   float64   `json:"points" gorm:"not null"`
	Orders   []Order   `json:"orders"`
	Receipts []Receipt `json:"receipts"`
}

func NewUser(name string, email string, password string, avatar string) (*User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
	}, nil
}

// UserService contains all `User`
// related business logic
type UserService interface {
	User(id uint) (*User, error)
	Users() ([]User, error)
	Create(*User) error
	Update(*User) error
	Delete(*User) error
}
