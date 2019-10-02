package aumo

import "github.com/jinzhu/gorm"

type Config struct {
	DB *gorm.DB
}

type Aumo struct {
	config Config
}

func New(c Config) Aumo {
	return Aumo{config: c}
}
