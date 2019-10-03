package aumo

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Points   float64
}

// CreateUser creates a user
func (a *Aumo) CreateUser(name, email, password string) (User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return User{}, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
		Points:   0,
	}

	a.DB.Create(user)

	return *user, nil
}

// GetUserByEmail returns a user that has a matching email
func (a *Aumo) GetUserByEmail(email string) (User, error) {
	var user User

	err := a.DB.First(&user, "email = ?", email)

	if err != nil {
		return User{}, nil
	}

	return user, nil
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return true
	}

	return false
}
