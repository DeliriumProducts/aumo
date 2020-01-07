package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/stretchr/testify/assert"
	"upper.io/db.v3"
)

func TestProductService(t *testing.T) {
	sess, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(sess)
		sess.Close()
	}()

	ps := mysql.NewProductService(sess)

	t.Run("create_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("TV", 500, "image.com", "ok", 5)
		err := ps.Create(pd)
		assert.Nil(t, err, "shouldn't return an error")

		pm := aumo.Product{}
		err = sess.Collection("products").Find("id", pd.ID).One(&pm)
		assert.Nil(t, err, "shouldn't return an error")
		assert.Equal(t, pm, *pd)
	})

	t.Run("get_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("Laptop", 100, "image.com", "it's a good laptop", 5)
		err := ps.Create(pd)
		assert.Nil(t, err, "shouldn't return an error")

		pm, err := ps.Product(pd.ID)
		assert.Nil(t, err, "shouldn't return an error")
		assert.Equal(t, *pd, *pm)
	})

	t.Run("get_products", func(t *testing.T) {
		defer TidyDB(sess)

		pds := []*aumo.Product{
			aumo.NewProduct("Phone", 100, "image.com", "it's a good phone", 5),
			aumo.NewProduct("ok", 100, "image.com", "it's a good phone", 5),
			aumo.NewProduct("TV", 100, "image.tv", "it's a good tv", 6),
		}

		for _, pd := range pds {
			err := ps.Create(pd)
			assert.Nil(t, err, "shouldn't return an error")
		}

		pms, err := ps.Products()
		assert.Nil(t, err, "it shouldn't return an error")
		assert.Equal(t, len(pms), len(pds), "it should have equal length")

		for i := 0; i < len(pms); i++ {
			assert.Equal(t, *pds[i], pms[i], "it should be equal")
		}
	})

	t.Run("delete_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		err := ps.Create(pd)
		assert.Nil(t, err, "shouldn't return an error")

		err = ps.Delete(pd.ID)
		assert.Nil(t, err, "shouldn't return an error")

		_, err = ps.Product(pd.ID)
		assert.Equal(t, err, db.ErrNoMoreRows)
	})

	t.Run("update_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("Computer", 400, "computer.com", "it's powerful", 10)
		err := ps.Create(pd)
		assert.Nil(t, err, "shouldn't return an error")

		pd.Name = "not a computer"
		err = ps.Update(pd.ID, pd)
		assert.Nil(t, err, "shouldn't return an error")

		pm, err := ps.Product(pd.ID)
		assert.Nil(t, err, "shouldn't return an error")
		assert.Equal(t, *pd, *pm)
	})
}
