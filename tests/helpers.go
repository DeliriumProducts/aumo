package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, us aumo.UserService) *aumo.User {
	u, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
	require.Nil(t, err, "shouldn't return an error")

	err = us.Create(u)
	require.Nil(t, err, "shouldn't return an error")

	return u
}

func createReceipt(t *testing.T, rs aumo.ReceiptService) *aumo.Receipt {
	r := aumo.NewReceipt(faker.AmountWithCurrency())

	err := rs.Create(r)
	require.Nil(t, err, "shouldn't return an error")

	return r
}

func createProduct(t *testing.T, ps aumo.ProductService, price float64, stock uint) *aumo.Product {
	p := aumo.NewProduct(faker.Word(), price, faker.URL(), faker.Sentence(), stock)

	err := ps.Create(p)
	require.Nil(t, err, "shouldn't return an error")

	return p
}
