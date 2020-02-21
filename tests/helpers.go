package tests

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo"
	"github.com/go-redis/redis/v7"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T, us aumo.UserStore) *aumo.User {
	u, err := aumo.NewUser(faker.FirstName(), faker.Email(), faker.Password(), faker.URL())
	require.Nil(t, err, "shouldn't return an error")

	err = us.Save(nil, u)
	require.Nil(t, err, "shouldn't return an error")

	return u
}

func createShop(t *testing.T, ss aumo.ShopStore) *aumo.Shop {
	s := aumo.NewShop(faker.Name(), faker.URL())

	err := ss.Save(nil, s)
	require.Nil(t, err, "shouldn't return an error")

	return s
}

func createReceipt(t *testing.T, rs aumo.ReceiptStore, s *aumo.Shop) *aumo.Receipt {
	r := aumo.NewReceipt(faker.AmountWithCurrency(), s.ID, 500)

	err := rs.Save(nil, r)
	require.Nil(t, err, "shouldn't return an error")

	return r
}

func createProduct(t *testing.T, ps aumo.ProductStore, s *aumo.Shop, price float64, stock uint) *aumo.Product {
	p := aumo.NewProduct(faker.Word(), price, faker.URL(), faker.Sentence(), stock, s.ID)

	err := ps.Save(nil, p)
	require.Nil(t, err, "shouldn't return an error")

	return p
}

func createSession(t *testing.T, r *redis.Client, user *aumo.User, expiryTime time.Duration) string {
	sID := faker.UUIDDigit()

	err := r.Set(sID, user.ID.String(), expiryTime).Err()
	require.Nil(t, err, "shouldn't return an error")
	require.NotEmpty(t, sID, "should return a session ID")

	return sID
}
