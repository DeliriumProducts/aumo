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
		err = sess.Collection("users").Find("id", u.ID).One(um)
		assert.Nil(t, err, "shouldn't return an error")
		um.Receipts = []aumo.Receipt{}
		um.Orders = []aumo.Order{}
		assert.Equal(t, *u, *um)
	})
}
