package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	ostore := mysql.NewOrderStore(sess)

	os := ordering.New(ostore, pstore, ustore)

	t.Run("place_order", func(t *testing.T) {
		defer TidyDB(sess)

		user := createUser(t, ustore)

		var price float64 = 500
		product := createProduct(t, pstore, price, 1)

		t.Run("valid", func(t *testing.T) {
			// Place order
			order, err := os.PlaceOrder(user.ID, product.ID)
			assert.Nil(t, err, "shouldn't return an error")

			// Update stock
			product.Stock--

			// Update points
			user.Points -= product.Price

			// Get product
			gotProduct, err := pstore.FindByID(nil, product.ID)
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, product.Stock, gotProduct.Stock, "should've decremented stock")

			// Get User
			gotUser, err := ustore.FindByID(nil, user.ID, true)
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, aumo.UserStartingPoints-price, gotUser.Points, "should've decremented user's points")

			// Check if order is in User's orders
			require.Contains(t, gotUser.Orders, *order)
		})

		t.Run("not_valid", func(t *testing.T) {
			// Place order
			order, err := os.PlaceOrder(user.ID, product.ID)
			require.Nil(t, order, "shouldn't have returned an error")
			require.NotNil(t, err, "should've returned an error")

			// Get product
			gotProduct, err := pstore.FindByID(nil, product.ID)
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, product.Stock, gotProduct.Stock, "shouldn't have been decremented")

			// Get user
			gotUser, err := ustore.FindByID(nil, user.ID, true)
			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, user.Points, gotUser.Points, "user shouldn't have been taxed")

			// Check if order isn't in User's orders
			require.Len(t, gotUser.Orders, 1)
		})
	})

}
