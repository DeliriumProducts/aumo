package tests

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
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
		pd := aumo.NewProduct("TV", 500, "image.com", "ok", 5)
		err := ps.Create(pd)
		spew.Dump(pd)
		assert.Nil(t, err, "didn't return an error")
		var pm aumo.Product
		db.Collection("products").Find("id", pd.ID).One(&pm)
		assert.Equal(t, pm, *pd)
	})
}
