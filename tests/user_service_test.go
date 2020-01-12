package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
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

	rs := mysql.NewReceiptService(sess)
	ps := mysql.NewProductService(sess)
	os := mysql.NewOrderService(sess)
	us := mysql.NewUserService(sess, rs, ps, os)

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
			err := rs.Create(r)
			assert.Nil(t, err, "shouldn't return an error")

			err = us.ClaimReceipt(u, r.ReceiptID)
			assert.Nil(t, err, "shouldn't return an error")

			p := aumo.NewProduct("TV", 500, "image.com", "it's good", 5)
			err = ps.Create(p)
			assert.Nil(t, err, "shouldn't return an error")

			err = us.PlaceOrder(u, p.ID)
			assert.Nil(t, err, "shouldn't return an error")

			um, err := us.User(u.ID, true)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, *u, *um, "should be equal")
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
			err = us.PlaceOrder(u, p.ID)
			assert.Nil(t, err, "shouldn't return an error")

			pm, err := ps.Product(p.ID)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, p.Stock-1, pm.Stock)

			assert.Equal(t, aumo.UserStartingPoints-price, u.Points)
		})
	})

	t.Run("claim_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		u, err := aumo.NewUser("Adrian", "adrian@pesho.com", "123456", "ok")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		t.Run("valid", func(t *testing.T) {
			r := aumo.NewReceipt("Paconi: 250LV")
			err := rs.Create(r)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, false, r.IsClaimed())

			err = us.ClaimReceipt(u, r.ReceiptID)
			assert.Nil(t, err, "shouldn't return an error")

			err = sess.Collection(mysql.ReceiptTable).Find("receipt_id", r.ReceiptID).One(r)
			assert.Nil(t, err, "shouldn't return an error")
			assert.Equal(t, true, r.IsClaimed())

			// TODO: fetch user, check if orders.Contains(order)
			// um, err := us.User(u.ID, true)
			// assert.Nil(t, err, "shouldn't return an error")
		})

		// t.Run("race_condition", func(t *testing.T) {
		// 	r := aumo.NewReceipt(u.ID, "Paconi: 250LV")
		// 	err := rs.Create(r)
		// 	assert.Nil(t, err, "shouldn't return an error")

		// })
	})
}
