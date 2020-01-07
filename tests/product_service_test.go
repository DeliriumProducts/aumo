package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/stretchr/testify/assert"
)

func TestProductService(t *testing.T) {
	db, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(db)
		db.Close()
	}()

	ps := mysql.NewProductService(db)

	t.Run("create_product", func(t *testing.T) {
		defer TidyDB(db)
		pd := aumo.NewProduct("TV", 500, "image.com", "ok", 5)
		err := ps.Create(pd)
		assert.Nil(t, err, "didn't return an error")
		var pm aumo.Product
		db.Collection("products").Find("id", pd.ID).One(&pm)
		assert.Equal(t, pm, *pd)
	})

	t.Run("get_product", func(t *testing.T) {
		defer TidyDB(db)
		pd := aumo.NewProduct("Laptop", 100, "image.com", "it's a good laptop", 5)
		ps.Create(pd)
		_, err := ps.Product(pd.ID)
		assert.Nil(t, err, "didn't return an error")
	})
}
