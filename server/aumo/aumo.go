// Package aumo provides business logic and necessary DB operations for the aumo project.
// All (a *Aumo) methods handle DB operations
// All (u *User) methods and any other struct methods are just
// convenient wrappers that simply mutate the struct or do some checking
// they DO NOT have any side effects (other than possibly mutating the original struct)
package aumo

import (
	"github.com/fr3fou/aumo/server/aumo/models"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB *gorm.DB
}

type Aumo struct {
	c  Config
	db *gorm.DB
}

func New(c Config) *Aumo {
	if c.DB == nil {
		panic("aumo: no db instance provided")
	}

	c.DB.AutoMigrate(&models.User{}, &models.ShopItem{}, &models.Receipt{})

	return &Aumo{
		c:  c,
		db: c.DB,
	}
}

// updateX is an internal helper function to update any struct
func (a *Aumo) updateX(i interface{}) error {
	return a.db.Model(i).Updates(i).Error
}

// deleteX is an internal helper function to update any struct
func (a *Aumo) deleteX(i interface{}) error {
	return a.db.Delete(i).Error
}

// firstX is an internal helper function to get the first row of any struct
func (a *Aumo) firstX(dest interface{}, where ...interface{}) error {
	return a.db.Set("gorm:auto_preload", true).First(dest, where...).Error
}

// findX is an internal helper function to get all of the rows of any struct
func (a *Aumo) findX(dest interface{}, where ...interface{}) error {
	return a.db.Set("gorm:auto_preload", true).Find(dest, where...).Error
}
