package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/deliriumproducts/aumo/products"
	"github.com/deliriumproducts/aumo/users"
	"github.com/stretchr/testify/assert"
)

func TestOrderService(t *testing.T) {
	sess, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(sess)
		sess.Close()
	}()

	pstore := mysql.NewProductStore(sess)
	ustore := mysql.NewUserStore(sess)

	os := ordering.New(mysql.NewOrderStore(sess), pstore, ustore)
	ps := products.New(pstore)
	us := users.New(ustore)

	t.Run("place_order", func(t *testing.T) {
		defer TidyDB(sess)

		user := createUser(t, us)

		var price float64 = 500
		product := createProduct(t, ps, price, 1)

		t.Run("valid", func(t *testing.T) {
			// Place order
			_, err := os.PlaceOrder(user.ID, product.ID)
			assert.Nil(t, err, "shouldn't return an error")

			// Update stock
			product.Stock--

			// Update points
			user.Points -= product.Price

			// Get product
			gotProduct, err := pstore.FindByID(nil, product.ID)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, product.Stock, gotProduct.Stock)

			// Get User
			gotUser, err := ustore.FindByID(nil, user.ID, false)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, aumo.UserStartingPoints-price, gotUser.Points)
		})

		t.Run("not_valid", func(t *testing.T) {
			// Place order
			_, err := os.PlaceOrder(user.ID, product.ID)
			assert.Equal(t, aumo.ErrNotInStock, err)

			// Get product
			gotProduct, err := pstore.FindByID(nil, product.ID)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, product.Stock, gotProduct.Stock, "shouldn't have been decremented")

			// Get user
			gotUser, err := ustore.FindByID(nil, user.ID, false)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, user.Points, gotUser.Points, "user shouldn't have been taxed")
		})
	})

}
