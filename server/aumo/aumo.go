package aumo

import "github.com/jinzhu/gorm"

type Config struct {
	DB *gorm.DB
}

type Aumo struct {
	Config
}

func New(c Config) Aumo {
	return Aumo{Config: c}
}
