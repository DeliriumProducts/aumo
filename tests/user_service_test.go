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

		u, err := aumo.NewUser("George", "go@sho.com", "1234", "asdf")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		um := &aumo.User{}
		err = sess.Collection(mysql.UserTable).Find("id", u.ID).One(um)
		assert.Nil(t, err, "shouldn't return an error")
		um.Receipts = []aumo.Receipt{}
		um.Orders = []aumo.Order{}
		assert.Equal(t, *u, *um)
	})

	t.Run("get_user", func(t *testing.T) {
		t.Run("by_id", func(t *testing.T) {
			defer TidyDB(sess)
			u, err := aumo.NewUser("George", "go@sho.com", "1234", "asdf")
			assert.Nil(t, err, "shouldn't return an error")
			err = us.Create(u)
			assert.Nil(t, err, "shouldn't return an error")

			t.Run("no_relations", func(t *testing.T) {
				us, err := us.User(u.ID, false)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *u, *us, "should be equal")
			})

			t.Run("with_relations", func(t *testing.T) {
				r := aumo.NewReceipt("Paconi: 250LV")
				err = rs.Create(r)
				assert.Nil(t, err, "shouldn't return an error")
				rc, err := rs.ClaimReceipt(u.ID, r.ReceiptID)
				assert.Nil(t, err, "shouldn't return an error")

				u.Receipts = append(u.Receipts, *rc)

				p := aumo.NewProduct("TV", 500, "image.com", "it's good", 5)
				err = ps.Create(p)
				assert.Nil(t, err, "shouldn't return an error")

				order, err := os.PlaceOrder(u.ID, p.ID)
				assert.Nil(t, err, "shouldn't return an error")

				u.Orders = append(u.Orders, *order)

				um, err := us.User(u.ID, true)

				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *u, *um, "should be equal")
			})
		})

		t.Run("by_email", func(t *testing.T) {
			defer TidyDB(sess)
			u, err := aumo.NewUser("George", "go@sho.com", "1234", "asdf")
			assert.Nil(t, err, "shouldn't return an error")
			err = us.Create(u)
			assert.Nil(t, err, "shouldn't return an error")

			t.Run("no_relations", func(t *testing.T) {
				us, err := us.User(u.ID, false)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *u, *us, "should be equal")
			})

			t.Run("with_relations", func(t *testing.T) {
				r := aumo.NewReceipt("Paconi: 250LV")
				err = rs.Create(r)
				assert.Nil(t, err, "shouldn't return an error")
				rc, err := rs.ClaimReceipt(u.ID, r.ReceiptID)
				assert.Nil(t, err, "shouldn't return an error")

				u.Receipts = append(u.Receipts, *rc)

				p := aumo.NewProduct("TV", 500, "image.com", "it's good", 5)
				err = ps.Create(p)
				assert.Nil(t, err, "shouldn't return an error")

				order, err := os.PlaceOrder(u.ID, p.ID)
				assert.Nil(t, err, "shouldn't return an error")

				u.Orders = append(u.Orders, *order)

				um, err := us.UserByEmail(u.Email, true)
				assert.Nil(t, err, "shouldn't return an error")
				assert.Equal(t, *u, *um, "should be equal")
			})
		})
	})

	t.Run("place_order", func(t *testing.T) {
		defer TidyDB(sess)

		u, err := aumo.NewUser("Jordan", "jord@an.com", "asdfjkl", "imgur.com")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		var price float64 = 500
		p := aumo.NewProduct("TV", price, "image.com", "it's good", 5)
		err = ps.Create(p)
		assert.Nil(t, err, "shouldn't return an error")

		t.Run("valid", func(t *testing.T) {
			_, err := os.PlaceOrder(u.ID, p.ID)
			assert.Nil(t, err, "shouldn't return an error")

			pm, err := ps.Product(p.ID)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, p.Stock-1, pm.Stock)

			us, err := us.User(u.ID, false)
			assert.Nil(t, err, "shouldn't return an error")

			assert.Equal(t, aumo.UserStartingPoints-price, us.Points)
		})
	})
}
