package tests

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, us aumo.UserStore) *aumo.User {
	u, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
	require.Nil(t, err, "shouldn't return an error")

	err = us.Save(nil, u)
	require.Nil(t, err, "shouldn't return an error")

	return u
}

func createReceipt(t *testing.T, rs aumo.ReceiptStore) *aumo.Receipt {
	r := aumo.NewReceipt(faker.AmountWithCurrency())

	err := rs.Save(nil, r)
	require.Nil(t, err, "shouldn't return an error")

	return r
}

func createProduct(t *testing.T, ps aumo.ProductStore, price float64, stock uint) *aumo.Product {
	p := aumo.NewProduct(faker.Word(), price, faker.URL(), faker.Sentence(), stock)

	err := ps.Save(nil, p)
	require.Nil(t, err, "shouldn't return an error")

	return p
}

func createSession(t *testing.T, r redis.Conn, user *aumo.User, expiryTime int) string {
	sID := faker.UUIDDigit()

	_, err := r.Do("SETEX", sID, expiryTime, user.ID)
	require.Nil(t, err, "shouldn't return an error")
	require.NotEmpty(t, sID, "should return a session ID")

	return sID
}
