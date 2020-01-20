package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
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

		product := aumo.NewProduct(faker.Word(), 500, faker.URL(), faker.Sentence(), 5)

		err := ps.Create(product)
		require.Nil(t, err, "shouldn't return an error")

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

		products := []*aumo.Product{
			aumo.NewProduct(faker.Word(), 1000, faker.URL(), faker.Sentence(), 99),
			aumo.NewProduct(faker.Word(), 20, faker.URL(), faker.Sentence(), 10),
			aumo.NewProduct(faker.Word(), 5000, faker.URL(), faker.Sentence(), 2),
		}

		for _, product := range products {
			err := pstore.Save(nil, product)
			require.Nil(t, err, "shouldn't return an error")
		}

		gotProducts, err := ps.Products()
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, len(gotProducts), len(products), "should have equal length")

		for i := 0; i < len(gotProducts); i++ {
			require.Equal(t, *products[i], gotProducts[i], "should be equal")
		}
	})

	t.Run("update_product", func(t *testing.T) {
		defer TidyDB(sess)

		product := createProduct(t, pstore, 500, 5)
		product.Name = "not a computer"

		err = ps.Update(product.ID, product)
		require.Nil(t, err, "shouldn't return an error")

		gotProduct, err := pstore.FindByID(nil, product.ID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, *product, *gotProduct)
	})

	t.Run("delete_product", func(t *testing.T) {
		defer TidyDB(sess)

		product := createProduct(t, pstore, 500, 5)

		err = ps.Delete(product.ID)
		require.Nil(t, err, "shouldn't return an error")

		_, err = pstore.FindByID(nil, product.ID)
		require.Equal(t, err, db.ErrNoMoreRows)
	})
}
