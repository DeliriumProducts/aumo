package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/deliriumproducts/aumo/products"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/deliriumproducts/aumo/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService(t *testing.T) {
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
	rs := receipt.New(mysql.NewReceiptStore(sess))

	t.Run("create_user", func(t *testing.T) {
		defer TidyDB(sess)

		user := createUser(t, us)

		gotUser, err := ustore.FindByID(nil, user.ID, false)

		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser)
	})

	testUserFetcher := func(t *testing.T, userFetcher func(user *aumo.User, relations bool) (*aumo.User, error)) {
		t.Run("no_relations", func(t *testing.T) {
			defer TidyDB(sess)
			user := createUser(t, us)
			gotUser, err := userFetcher(user, false)

			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, *user, *gotUser, "should be equal")
		})

		t.Run("with_relations", func(t *testing.T) {
			t.Run("empty_relations", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, us)

				// Get the user
				gotUser, err := userFetcher(user, true)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *user, *gotUser, "should be equal")
			})
			t.Run("only_receipts", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, us)

				// Create a receipt
				receipt := createReceipt(t, rs)

				var err error

				// Claim the receipt
				receipt, err = rs.ClaimReceipt(user.ID, receipt.ReceiptID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the receipt
				user.Receipts = append(user.Receipts, *receipt)

				// Get the user
				gotUser, err := userFetcher(user, true)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *user, *gotUser, "should be equal")
			})
			t.Run("only_orders", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, us)

				// Create a product
				product := createProduct(t, ps, 500, 5)

				// Buy the product
				order, err := os.PlaceOrder(user.ID, product.ID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the order
				user.Orders = append(user.Orders, *order)

				// Subtract points
				user.Points -= product.Price

				// Get the user
				gotUser, err := userFetcher(user, true)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *user, *gotUser, "should be equal")

			})
			t.Run("all_relations", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, us)

				// Create a receipt
				receipt := createReceipt(t, rs)

				var err error

				// Claim the receipt
				receipt, err = rs.ClaimReceipt(user.ID, receipt.ReceiptID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the receipt
				user.Receipts = append(user.Receipts, *receipt)

				// Create a product
				product := createProduct(t, ps, 500, 5)

				// Buy the product
				order, err := os.PlaceOrder(user.ID, product.ID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the order
				user.Orders = append(user.Orders, *order)

				// Subtract points
				user.Points -= product.Price

				// Get the user
				gotUser, err := userFetcher(user, true)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *user, *gotUser, "should be equal")
			})
		})
	}

	t.Run("get_user", func(t *testing.T) {
		t.Run("by_id", func(t *testing.T) {
			testUserFetcher(t, func(u *aumo.User, relations bool) (*aumo.User, error) {
				return us.User(u.ID, relations)
			})
		})

		t.Run("by_email", func(t *testing.T) {
			testUserFetcher(t, func(u *aumo.User, relations bool) (*aumo.User, error) {
				return us.UserByEmail(u.Email, relations)
			})
		})
	})
}
