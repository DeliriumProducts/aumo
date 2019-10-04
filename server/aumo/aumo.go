// aumo is the internal library that contains all
// business logic and necessary DB operations for the aumo project.
// All (a *Aumo) methods handle DB operations
// All (u *User) methods and any other struct methods are just
// convenient wrappers that simply mutate the struct or do some checking
// they DO NOT have any side effects (other than possibly mutating the original struct)
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

	c.DB.AutoMigrate(&User{}, &ShopItem{}, &Receipt{})

	return Aumo{
		Config: c,
	}
}

// updateX is an internal helper function to update any struct
func (a *Aumo) updateX(i interface{}) error {
	return a.DB.Model(i).Updates(i).Error
}

// deleteX is an internal helper function to update any struct
func (a *Aumo) deleteX(i interface{}) error {
	return a.DB.Delete(i).Error
}

// firstX is an internal helper function to get the first row of any struct
func (a *Aumo) firstX(dest interface{}, where ...interface{}) error {
	return a.DB.First(dest, where...).Error
}

// findX is an internal helper function to get all of the rows of any struct
func (a *Aumo) findX(dest interface{}, where ...interface{}) error {
	return a.DB.Find(dest, where...).Error
}
