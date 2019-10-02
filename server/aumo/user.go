package aumo

import "github.com/jihzu/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Points   float64
}
