package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/products"
	"github.com/stretchr/testify/require"
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
	pstore := mysql.NewProductStore(sess)
	ps := products.New(pstore)

	t.Run("create_product", func(t *testing.T) {
		defer TidyDB(sess)

		product := createProduct(t, pstore, 500, 5)

		gotProduct, err := pstore.FindByID(nil, product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
	})

	t.Run("get_product", func(t *testing.T) {
		defer TidyDB(sess)

		product := createProduct(t, pstore, 500, 5)

		gotProduct, err := ps.Product(product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
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
			require.Nil(t, err, "shouldn't return an error")
		}

		pms, err := ps.Products()
		require.Nil(t, err, "it shouldn't return an error")
		require.Equal(t, len(pms), len(pds), "it should have equal length")

		for i := 0; i < len(pms); i++ {
			require.Equal(t, *pds[i], pms[i], "it should be equal")
		}
	})

	t.Run("delete_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("Phone", 99, "image.com", "it's a good phone", 2)
		err := ps.Create(pd)
		require.Nil(t, err, "shouldn't return an error")

		err = ps.Delete(pd.ID)
		require.Nil(t, err, "shouldn't return an error")

		_, err = ps.Product(pd.ID)
		require.Equal(t, err, db.ErrNoMoreRows)
	})

	t.Run("update_product", func(t *testing.T) {
		defer TidyDB(sess)

		pd := aumo.NewProduct("Computer", 400, "computer.com", "it's powerful", 10)
		err := ps.Create(pd)
		require.Nil(t, err, "shouldn't return an error")

		pd.Name = "not a computer"
		err = ps.Update(pd.ID, pd)
		require.Nil(t, err, "shouldn't return an error")

		pm, err := ps.Product(pd.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *pd, *pm)
	})
}
