package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/users"
	"github.com/stretchr/testify/require"
)

func TestInitialAdmin(t *testing.T) {
	sess, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(sess)
		sess.Close()
	}()

	ustore := mysql.NewUserStore(sess)

	t.Run("doesn't_exist", func(t *testing.T) {
		defer TidyDB(sess)

		// Arrange
		password := "123456"
		email := "admin@deliriumproducts.me"

		// Act
		user, err := users.InitialAdmin(ustore, password, email)
		require.Nil(t, err, "shouldn't return an error")
		require.NotNil(t, user, "should not be nil")

		// Assert
		gotUser, err := ustore.FindByEmail(nil, email, false)
		require.Nil(t, err, "shouldn't return an error")
		require.NotNil(t, *gotUser, "should not be nil")
		require.Equal(t, aumo.Admin, gotUser.Role, "should be an admin")
		require.True(t, gotUser.IsVerified, "should be verified")
	})

	t.Run("already_exists", func(t *testing.T) {
		t.Run("is_admin", func(t *testing.T) {
			defer TidyDB(sess)
			// Arrange
			password := "123456"
			email := "admin@deliriumproducts.me"
			user, err := users.InitialAdmin(ustore, password, email)
			require.Nil(t, err, "shouldn't return an error")
			require.NotNil(t, user, "should not be nil")

			// Act
			user1, err := users.InitialAdmin(ustore, password, email)
			require.Nil(t, err, "shouldn't return an error")
			require.NotNil(t, user, "should not be nil")
			require.Equal(t, user, user1, "should be equal")

			// Assert
			gotUsers, err := ustore.FindAll(nil)
			require.Nil(t, err, "shouldn't return an error")
			require.Len(t, gotUsers, 1, "should only have the first admin")
			require.True(t, gotUsers[0].IsVerified, "should be verified")
			require.Equal(t, aumo.Admin, gotUsers[0].Role, "should be an admin")
		})

		t.Run("isn't_admin", func(t *testing.T) {
			defer TidyDB(sess)

			// Create Admin
			password := "123456"
			email := "admin@deliriumproducts.me"
			user, err := users.InitialAdmin(ustore, password, email)
			require.Nil(t, err, "shouldn't return an error")
			require.NotNil(t, user, "should not be nil")

			// Update role
			user.Role = aumo.Customer
			err = ustore.Update(nil, user.ID.String(), user)
			require.Nil(t, err, "shouldn't return an error")

			// Act
			user, err = users.InitialAdmin(ustore, password, email)
			require.Nil(t, err, "shouldn't return an error")
			require.NotNil(t, user, "should not be nil")

			// Assert
			gotUsers, err := ustore.FindAll(nil)
			require.Nil(t, err, "shouldn't return an error")
			require.Len(t, gotUsers, 1, "should only have the first admin")
			require.True(t, gotUsers[0].IsVerified, "should be verified")
			require.Equal(t, aumo.Admin, gotUsers[0].Role, "should be an admin")
		})
	})
}
