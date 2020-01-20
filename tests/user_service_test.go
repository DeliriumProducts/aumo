package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/ordering"
	"github.com/deliriumproducts/aumo/receipt"
	"github.com/deliriumproducts/aumo/users"
	"github.com/stretchr/testify/require"
	"upper.io/db.v3"
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
	ostore := mysql.NewOrderStore(sess)
	rstore := mysql.NewReceiptStore(sess)

	os := ordering.New(ostore, pstore, ustore)
	us := users.New(ustore)
	rs := receipt.New(rstore, ustore)

	t.Run("create_user", func(t *testing.T) {
		defer TidyDB(sess)

		user, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
		require.Nil(t, err, "shouldn't return an error")

		err = us.Create(user)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID, false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser)
	})

	testUserFetcher := func(t *testing.T, userFetcher func(user *aumo.User, relations bool) (*aumo.User, error)) {
		t.Run("no_relations", func(t *testing.T) {
			defer TidyDB(sess)
			user := createUser(t, ustore)
			gotUser, err := userFetcher(user, false)

			require.Nil(t, err, "shouldn't return an error")
			require.Equal(t, *user, *gotUser, "should be equal")
		})

		t.Run("with_relations", func(t *testing.T) {
			t.Run("empty_relations", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, ustore)

				// Get the user
				gotUser, err := userFetcher(user, true)
				require.Nil(t, err, "shouldn't return an error")
				require.Equal(t, *user, *gotUser, "should be equal")
			})
			t.Run("only_receipts", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, ustore)

				// Create a receipt
				receipt := createReceipt(t, rstore)

				var err error

				// Claim the receipt
				receipt, err = rs.ClaimReceipt(user.ID, receipt.ReceiptID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the receipt
				user.Receipts = append(user.Receipts, *receipt)

				// Add points
				user.Points += aumo.UserPointsPerReceipt

				// Get the user
				gotUser, err := userFetcher(user, true)
				require.Nil(t, err, "shouldn't return an error")
				require.Equal(t, *user, *gotUser, "should be equal")
			})
			t.Run("only_orders", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, ustore)

				// Create a product
				product := createProduct(t, pstore, 500, 5)

				// Buy the product
				order, err := os.PlaceOrder(user.ID, product.ID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the order
				user.Orders = append(user.Orders, *order)

				// Subtract points
				user.Points -= product.Price

				// Get the user
				gotUser, err := userFetcher(user, true)
				require.Nil(t, err, "shouldn't return an error")
				require.Equal(t, *user, *gotUser, "should be equal")

			})
			t.Run("all_relations", func(t *testing.T) {
				defer TidyDB(sess)
				user := createUser(t, ustore)

				// Create a receipt
				receipt := createReceipt(t, rstore)

				var err error

				// Claim the receipt
				receipt, err = rs.ClaimReceipt(user.ID, receipt.ReceiptID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the receipt
				user.Receipts = append(user.Receipts, *receipt)
				// Add points
				user.Points += aumo.UserPointsPerReceipt

				// Create a product
				product := createProduct(t, pstore, 500, 5)

				// Buy the product
				order, err := os.PlaceOrder(user.ID, product.ID)
				require.Nil(t, err, "shouldn't return an error")

				// Add the order
				user.Orders = append(user.Orders, *order)

				// Subtract points
				user.Points -= product.Price

				// Get the user
				gotUser, err := userFetcher(user, true)
				require.Nil(t, err, "shouldn't return an error")
				require.Equal(t, *user, *gotUser, "should be equal")
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

	t.Run("update_user", func(t *testing.T) {
		user := createUser(t, ustore)
		user.Name = "New Name"

		err := us.Update(user.ID, user)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID, false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser, "should be equal")
	})

	t.Run("delete_user", func(t *testing.T) {
		user := createUser(t, ustore)
		user.Name = "New Name"

		err := us.Delete(user.ID)
		require.Nil(t, err, "shouldn't return an error")

		_, err = ustore.FindByID(nil, user.ID, false)
		require.Equal(t, err, db.ErrNoMoreRows)
	})
}
