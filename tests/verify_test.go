package tests

import (
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/deliriumproducts/aumo/mail"
	"github.com/deliriumproducts/aumo/verifications"
	"github.com/stretchr/testify/require"
)

func TestVerification(t *testing.T) {
	r, err := SetupRedis()
	if err != nil {
		t.Error(err)
	}

	defer TidyRedis(r)

	mailer := mail.DevMailer{}
	verifier := verifications.New(mailer, r)

	t.Run("first_part", func(t *testing.T) {
		defer TidyRedis(r)

		val := "foo"

		token, err := verifier.Send(faker.Email(), val, faker.Sentence(), faker.Sentence(), faker.URL(), 1*time.Hour)
		require.Nil(t, err, "shouldn't return an error")
		require.NotEmpty(t, token, "should return a token")

		gotVal, err := r.Get(token).Result()
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, val, gotVal, "should be equal")
	})

	t.Run("second_part", func(t *testing.T) {
		defer TidyRedis(r)

		val := "foo"
		key := faker.UUIDHyphenated()

		err := r.Set(key, val, 0).Err()
		require.Nil(t, err, "shouldn't return an error")

		gotVal, err := verifier.Verify(key)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, val, gotVal, "should be equal")

		gotVal, err = r.Get(key).Result()
		require.NotNil(t, err, "should return an error")
		require.Equal(t, "", gotVal, "should be empty")
	})

}
