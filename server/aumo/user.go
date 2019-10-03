package aumo

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"-"`
	Points   float64    `json:"points"`
	Orders   []ShopItem `gorm:"many2many:user_shop_item;"`
	a        *Aumo
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
		Orders:   []ShopItem{},
	}

	if err := a.DB.Create(user).Error; err != nil {
		return User{}, err
	}

	user.a = a

	return *user, nil
}

// GetUserByEmail returns a user that has a matching email
func (a *Aumo) GetUserByEmail(email string) (User, error) {
	return a.getUser("email = ?", email)
}

// GetUserById returns a user that has a matching id
func (a *Aumo) GetUserById(id uint) (User, error) {
	return a.getUser("id = ?", id)
}

// getUser is an internal helper function to quickly get a user
func (a *Aumo) getUser(where ...interface{}) (User, error) {
	var user User

	err := a.DB.First(&user, where...).Error

	if err != nil {
		return User{}, err
	}

	user.a = a

	return user, nil
}

// SetUserPoints sets the user's points to the provided ones
func (u *User) SetUserPoints(points float64) error {
	return u.a.DB.Model(u).Update("points", points).Error
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

func (u *User) BuyItem(si ShopItem, quantity uint) error {
	if u.Points-si.Price*float64(quantity) < 0 {
		return ErrNotSufficientPoints
	}

	if si.Quantity-quantity < 0 {
		return ErrNotInStock
	}

	err := u.SetUserPoints(u.Points - si.Price*float64(quantity))

	if err != nil {
		return err
	}

	err = si.SetQuantity(si.Quantity - quantity)
	if err != nil {
		return err
	}

	return u.a.DB.Model(u).Association("Orders").Append(si).Error
}
