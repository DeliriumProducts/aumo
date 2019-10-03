package aumo

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	ErrNotSufficientPoints = errors.New("aumo: user doesn't have enough points to buy this item")
	ErrNotInStock          = errors.New("aumo: shop doesn't have enough stock of the item")
)

type Config struct {
	DB *gorm.DB
}

type Aumo struct {
	Config
}

func New(c Config) Aumo {
	if c.DB == nil {
		panic("aumo: no db instance provided")
	}

	c.DB.AutoMigrate(&User{}, &ShopItem{})

	return Aumo{
		Config: c,
	}
}

// updateX is an internal helper function to update any struct
func (a *Aumo) updateX(old, new interface{}) error {
	return a.DB.Model(old).Updates(new).Error
}

// deleteX is an internal helper function to update any struct
func (a *Aumo) deleteX(i interface{}) error {
	return a.DB.Delete(i).Error
}
