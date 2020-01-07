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

	defer db.Close()

	ps := mysql.NewProductService(db)

	t.Run("create_product", func(t *testing.T) {
		pd := &aumo.Product{}
		err := ps.Create(pd)
		assert.Nil(t, err, "didn't return an error")
	})

}
