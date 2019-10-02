package aumo

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Points   float64
}
