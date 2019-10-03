package aumo

import (
	"github.com/jinzhu/gorm"
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
