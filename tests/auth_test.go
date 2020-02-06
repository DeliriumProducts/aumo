package tests

import (
	"testing"
	"time"

	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/stretchr/testify/require"
)

func TestAuthenticator(t *testing.T) {
	sess, err := SetupDB()
	if err != nil {
		t.Error(err)
	}

	r, err := SetupRedis()
	if err != nil {
		t.Error(err)
	}

	// Cleanup
	defer func() {
		TidyDB(sess)
		TidyRedis(r)
		sess.Close()
		r.Close()
	}()

	ustore := mysql.NewUserStore(sess)

	a := auth.New(r, ustore, "http://localhost:3000", "/", time.Hour*24)

	t.Run("new_session", func(t *testing.T) {
		defer TidyRedis(r)
		user := createUser(t, ustore)

		sID, err := a.NewSession(user)
		require.Nil(t, err, "shouldn't return an error")
		require.NotEmpty(t, sID, "should return a session ID")

		uID, err := r.Get(sID).Result()
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, user.ID.String(), uID)
	})

	t.Run("get_user_from_session", func(t *testing.T) {
		defer TidyRedis(r)

		user := createUser(t, ustore)
		sID := createSession(t, r, user, time.Hour*24)

		gotUser, err := a.Get(sID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, user.ID, gotUser.ID)
	})

	t.Run("delete_session", func(t *testing.T) {
		defer TidyRedis(r)

		user := createUser(t, ustore)
		sID := createSession(t, r, user, time.Hour*24)

		err = a.Del(sID)
		require.Nil(t, err, "shouldn't return an error")

		uID, err := r.Get(sID).Uint64()
		require.NotNil(t, err, "should return an error")
		require.Empty(t, uID, "shouldn't return a user ID")
	})
}
