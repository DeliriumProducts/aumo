package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/stretchr/testify/assert"
)

func TestUserPlaceOrder(t *testing.T) {
	t.Run("valid_purchase", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		assert.Nil(t, err, "shouldn't return an err")
		u.ID = 1

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		assert.Nil(t, err, "shouldn't return an err")

		assert.Contains(t, u.Orders, *o, "the order should be appeneded to the array")
	})

	t.Run("not_enough_points", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		assert.Nil(t, err, "shouldn't return an err")
		u.ID = 1
		u.Points = 0

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		assert.Equal(t, err, aumo.ErrNotSufficientPoints)

		assert.NotContains(t, u.Orders, *o, "the order shouldn't have been appended to the array")
	})

	t.Run("not_in_stock", func(t *testing.T) {
		u, err := aumo.NewUser("Gosho", "gosho@abv.bg", "123456", "pesho.com/gosho.png")
		assert.Nil(t, err, "shouldn't return an err")
		u.ID = 1

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
		u.ID = 1
		u.Points = 0

		p := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 0)
		p.ID = 1

		o := aumo.NewOrder(u, p)
		err = u.PlaceOrder(o)
		assert.NotNil(t, err, "should return an err")

		assert.NotContains(t, u.Orders, *o, "the order shouldn't have been appended to the array")
	})
}
