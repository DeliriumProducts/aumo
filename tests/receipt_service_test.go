package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/stretchr/testify/assert"
)

func TestReceiptService(t *testing.T) {
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

	t.Run("create_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		u, err := aumo.NewUser("George", "go@sho.com", "1234", "asdf")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		rp := aumo.NewReceipt(u.ID, "Paconi: 230")
		err = rs.Create(rp)
		assert.Nil(t, err, "shouldn't return an error")

		r := aumo.Receipt{}
		err = sess.Collection("receipts").Find("id", rp.ID).One(&r)
		assert.Nil(t, err, "shouldn't return an error")
		assert.Equal(t, r, *rp)
	})

	t.Run("get_receipt", func(t *testing.T) {
		defer TidyDB(sess)

		u, err := aumo.NewUser("Pesho", "pe@sho.com", "123123", "asdff")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		rp := aumo.NewReceipt(u.ID, "CBA: 30 Leva")
		err = rs.Create(rp)
		assert.Nil(t, err, "shouldn't return an error")

		r, err := rs.Receipt(rp.ID)
		assert.Nil(t, err, "shouldn't return an error")
		assert.Equal(t, *r, *rp)
	})

	t.Run("get_receipts", func(t *testing.T) {
		defer TidyDB(sess)

		u, err := aumo.NewUser("Pesho", "pe@sho.com", "123123", "asdff")
		assert.Nil(t, err, "shouldn't return an error")
		err = us.Create(u)
		assert.Nil(t, err, "shouldn't return an error")

		rds := []*aumo.Receipt{
			aumo.NewReceipt(u.ID, "CBA: 30 Leva"),
			aumo.NewReceipt(u.ID, "Pesho: 60 Leva"),
			aumo.NewReceipt(u.ID, "Pesho: 100 Leva"),
		}

		for _, rd := range rds {
			err := rs.Create(rd)
			assert.Nil(t, err, "shouldn't return an error")
		}

		rms, err := rs.Receipts()
		assert.Nil(t, err, "it shouldn't return an error")
		assert.Equal(t, len(rms), len(rds), "it should have equal length")

		for i := 0; i < len(rms); i++ {
			assert.Equal(t, *rds[i], rms[i], "it should be equal")
		}
	})
}
