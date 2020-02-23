package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/deliriumproducts/aumo/products"
	"github.com/stretchr/testify/require"
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
	sstore := mysql.NewShopStore(sess)
	ps := products.New(pstore)

	t.Run("create_product", func(t *testing.T) {
		defer TidyDB(sess)

		s := createShop(t, sstore)
		product := aumo.NewProduct(faker.Word(), 500, faker.URL(), faker.Sentence(), 50, s.ID)
		product.Shop = s

		err := ps.Create(product)
		require.Nil(t, err, "shouldn't return an error")

		gotProduct, err := pstore.FindByID(nil, product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
	})

	t.Run("get_product", func(t *testing.T) {
		defer TidyDB(sess)
		s := createShop(t, sstore)
		product := createProduct(t, pstore, s, 500, 5)
		product.Shop = s

		gotProduct, err := ps.Product(product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
	})

	t.Run("get_products", func(t *testing.T) {
		defer TidyDB(sess)

		products := []*aumo.Product{
			createProduct(t, pstore, createShop(t, sstore), 50, 18),
			createProduct(t, pstore, createShop(t, sstore), 34, 9),
			createProduct(t, pstore, createShop(t, sstore), 234, 20),
		}

		gotProducts, err := ps.Products()
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, len(products), len(gotProducts), "should have equal length")

		for i := 0; i < len(gotProducts); i++ {
			require.Equal(t, *products[i], gotProducts[i], "should be equal")
		}
	})

	t.Run("update_product", func(t *testing.T) {
		defer TidyDB(sess)

		s := createShop(t, sstore)
		product := createProduct(t, pstore, s, 500, 5)
		product.Shop = s
		product.Name = "not a computer"

		err = ps.Update(product.ID, product)
		require.Nil(t, err, "shouldn't return an error")

		gotProduct, err := pstore.FindByID(nil, product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
	})

	t.Run("delete_product", func(t *testing.T) {
		defer TidyDB(sess)

		product := createProduct(t, pstore, createShop(t, sstore), 500, 5)

		err = ps.Delete(product.ID)
		require.Nil(t, err, "shouldn't return an error")

		_, err = pstore.FindByID(nil, product.ID)
		require.Equal(t, err, aumo.ErrProductNotFound)
	})
}
