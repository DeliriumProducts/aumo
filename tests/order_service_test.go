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
