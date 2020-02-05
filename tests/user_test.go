package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Run("lowers_email", func(t *testing.T) {
		u, err := aumo.NewUser("Pesho", "FOO@bar.com", "123456", "pesho.com/pesho.png")
		require.Nil(t, err)
		require.Equal(t, u.Email, "foo@bar.com")
	})

	t.Run("trims_email", func(t *testing.T) {
		u, err := aumo.NewUser("Pesho", "           foo@bar.com            ", "123456", "pesho.com/pesho.png")
		require.Nil(t, err)
		require.Equal(t, u.Email, "foo@bar.com")
	})

	t.Run("trims_and_lowers_email", func(t *testing.T) {
		u, err := aumo.NewUser("Pesho", "           foo@BAR.com            ", "123456", "pesho.com/pesho.png")
		require.Nil(t, err)
		require.Equal(t, u.Email, "foo@bar.com")
	})
}

func TestUserPlaceOrder(t *testing.T) {
	t.Run("valid_purchase", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		require.Nil(t, err, "shouldn't return an err")

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		require.Nil(t, err, "shouldn't return an err")

		require.Contains(t, u.Orders, *o, "the order should be appeneded to the array")
	})

	t.Run("not_enough_points", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		require.Nil(t, err, "shouldn't return an err")
		u.Points = 0

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		require.Equal(t, err, aumo.ErrNotSufficientPoints)

		require.NotContains(t, u.Orders, *o, "the order shouldn't have been appended to the array")
	})

	t.Run("not_in_stock", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		assert.Nil(t, err, "shouldn't return an err")

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 0)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		assert.Equal(t, err, aumo.ErrNotInStock)

		assert.NotContains(t, u.Orders, *o, "the order shouldn't have been appended to the array")
	})

	t.Run("both_not_in_stock_and_not_enough_points", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		assert.Nil(t, err, "shouldn't return an err")
		u.Points = 0

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 0)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		assert.NotNil(t, err, "should return an err")

		assert.NotContains(t, u.Orders, *o, "the order shouldn't have been appended to the array")
	})
}
