package users

import (
	"errors"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3"
)

// InitialAdmin creates an Admin and returns it if it was created
func InitialAdmin(us aumo.UserStore, initialPassword, initialEmail string) (*aumo.User, error) {
	user, err := us.FindByEmail(nil, initialEmail, false)
	if errors.Is(err, db.ErrNoMoreRows) {
		// Doesn't exist
		user, err := aumo.NewUser("Aumo Admin", initialEmail, initialPassword, "https://i.imgur.com/QUEMEDP.png")
		if err != nil {
			return nil, err
		}

		user.Role = aumo.Admin
		user.IsVerified = true

		return user, us.Save(nil, user)
	}

	// Exists but wrong role
	if user.Role != aumo.Admin {
		user.Role = aumo.Admin
		return user, us.Update(nil, user.ID, user)
	}

	// Exists
	return user, err
}
