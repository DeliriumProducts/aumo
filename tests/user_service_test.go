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
	sstore := mysql.NewShopStore(sess)

	os := ordering.New(ostore, pstore, ustore)
	us := users.New(ustore)
	rs := receipt.New(rstore, ustore)

	t.Run("create_user", func(t *testing.T) {
		defer TidyDB(sess)

		user, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
		require.Nil(t, err, "shouldn't return an error")

		err = us.Create(user)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID.String(), false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser)
	})

	testUserFetcher := func(t *testing.T, userFetcher func(user *aumo.User, relations bool) (*aumo.User, error)) {
		// helper func for require.Nil() on the err
		user := func(t *testing.T) *aumo.User {
			u, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
			require.Nil(t, err, "shouldn't return an err")
			return u
		}

		tests := []struct {
			name      string
			user      *aumo.User
			receipts  []aumo.Receipt
			products  []aumo.Product
			relations bool
		}{
			{
				"no_relations",
				user(t),
				[]aumo.Receipt{},
				[]aumo.Product{},
				false,
			},
			{
				"one_order",
				user(t),
				[]aumo.Receipt{},
				[]aumo.Product{*createProduct(t, pstore, createShop(t, sstore), 500, 5)},
				true,
			},
			{
				"one_receipt",
				user(t),
				[]aumo.Receipt{*createReceipt(t, rstore, createShop(t, sstore))},
				[]aumo.Product{},
				true,
			},
			{
				"many_orders",
				user(t),
				[]aumo.Receipt{},
				[]aumo.Product{
					*createProduct(t, pstore, createShop(t, sstore), 250, 5),
					*createProduct(t, pstore, createShop(t, sstore), 1000, 1),
					*createProduct(t, pstore, createShop(t, sstore), 300, 8),
				},
				true,
			},
			{
				"many_receipts",
				user(t),
				[]aumo.Receipt{
					*createReceipt(t, rstore, createShop(t, sstore)),
					*createReceipt(t, rstore, createShop(t, sstore)),
				},
				[]aumo.Product{},
				true,
			},
			{
				"many_orders_many_receipts",
				user(t),
				[]aumo.Receipt{
					*createReceipt(t, rstore, createShop(t, sstore)),
					*createReceipt(t, rstore, createShop(t, sstore)),
				},
				[]aumo.Product{
					*createProduct(t, pstore, createShop(t, sstore), 100, 5),
					*createProduct(t, pstore, createShop(t, sstore), 500, 1),
					*createProduct(t, pstore, createShop(t, sstore), 75, 8),
				},
				true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				defer TidyDB(sess)

				err = ustore.Save(nil, tt.user)
				require.Nil(t, err, "shouldn't return an error")

				for _, receipt := range tt.receipts {
					r := receipt

					err = rstore.Save(nil, &r)
					require.Nil(t, err, "shouldn't return an error")

					receipt, err := rs.ClaimReceipt(tt.user.ID.String(), r.ReceiptID.String())
					require.Nil(t, err, "shouldn't return an error")

					tt.user.Points += aumo.UserPointsPerReceipt
					tt.user.Receipts = append(tt.user.Receipts, *receipt)
				}

				for _, product := range tt.products {
					p := product

					err = pstore.Save(nil, &p)
					require.Nil(t, err, "shouldn't return an error")

					order, err := os.PlaceOrder(tt.user.ID.String(), p.ID)
					require.Nil(t, err, "shouldn't return an error")

					tt.user.Points -= p.Price
					tt.user.Orders = append(tt.user.Orders, *order)
				}

				gotUser, err := userFetcher(tt.user, tt.relations)

				require.Nil(t, err, "shouldn't return an error")
				require.ElementsMatch(t, gotUser.Receipts, tt.user.Receipts, "should be equal")
				require.ElementsMatch(t, gotUser.Orders, tt.user.Orders, "should be equal")

				tt.user.Receipts = []aumo.Receipt{}
				tt.user.Orders = []aumo.Order{}

				gotUser.Receipts = []aumo.Receipt{}
				gotUser.Orders = []aumo.Order{}

				require.Equal(t, gotUser, tt.user, "should be equal")
			})
		}
	}

	t.Run("get_user", func(t *testing.T) {
		t.Run("by_id", func(t *testing.T) {
			testUserFetcher(t, func(u *aumo.User, relations bool) (*aumo.User, error) {
				return us.User(u.ID.String(), relations)
			})
		})

		t.Run("by_email", func(t *testing.T) {
			testUserFetcher(t, func(u *aumo.User, relations bool) (*aumo.User, error) {
				return us.UserByEmail(u.Email, relations)
			})
		})
	})

	t.Run("update_user", func(t *testing.T) {
		defer TidyDB(sess)
		user := createUser(t, ustore)
		user.Name = "New Name"

		err := us.Update(user.ID.String(), user)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID.String(), false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser, "should be equal")
	})

	t.Run("edit_role", func(t *testing.T) {
		defer TidyDB(sess)
		user := createUser(t, ustore)
		user.Role = aumo.Admin

		err := us.EditRole(user.ID.String(), aumo.Admin)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID.String(), false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser, "should be equal")
	})

	t.Run("add_points", func(t *testing.T) {
		defer TidyDB(sess)
		user := createUser(t, ustore)
		var points float64 = 500
		user.Points += points

		err := us.AddPoints(user.ID.String(), points)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID.String(), false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser, "should be equal")
	})

	t.Run("sub_points", func(t *testing.T) {
		defer TidyDB(sess)
		user := createUser(t, ustore)
		var points float64 = 500
		user.Points -= points

		err := us.SubPoints(user.ID.String(), points)
		require.Nil(t, err, "shouldn't return an error")

		gotUser, err := ustore.FindByID(nil, user.ID.String(), false)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *user, *gotUser, "should be equal")
	})

	t.Run("delete_user", func(t *testing.T) {
		defer TidyDB(sess)
		user := createUser(t, ustore)

		err := us.Delete(user.ID.String())
		require.Nil(t, err, "shouldn't return an error")

		_, err = ustore.FindByID(nil, user.ID.String(), false)
		require.Equal(t, err, aumo.ErrUserNotFound)
	})
}
