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

	return Aumo{
		Config: c,
	}
}
