package tests

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
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

		us, err := us.User(u.ID, false)
		assert.Nil(t, err, "shouldn't return an error")
		us.Receipts = []aumo.Receipt{}
		us.Orders = []aumo.Order{}
		assert.Equal(t, *u, *us, "should be equal")
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
			spew.Dump(r)
			assert.Nil(t, err, "shouldn't return an error")

			err = us.ClaimReceipt(u, r.ID)
			assert.Nil(t, err, "shouldn't return an error")
		})

		// t.Run("race_condition", func(t *testing.T) {
		// 	r := aumo.NewReceipt(u.ID, "Paconi: 250LV")
		// 	err := rs.Create(r)
		// 	assert.Nil(t, err, "shouldn't return an error")

		// })

	})
}
