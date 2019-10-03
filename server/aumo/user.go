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
	return a.getUser(&User{}, "email = ?", email)
}

// GetUserById returns a user that has a matching id
func (a *Aumo) GetUserById(id uint) (User, error) {
	return a.getUser(&User{}, "id = ?", id)
}

// getUser is an internal helper function to quickly get a user
func (a *Aumo) getUser(out interface{}, where ...interface{}) (User, error) {
	var user User

	err := a.DB.First(out, where...)

	if err != nil {
		return User{}, nil
	}

	return user, nil
}

// SetUserPoints sets the user's  points to the provided ones
func (a *Aumo) SetUserPoints(u *User, points float64) {
	a.DB.Model(u).Update("points", points)
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return true
	}

	return false
}
