package tests

import (
	"testing"

	"github.com/deliriumproducts/aumo/auth"
	"github.com/deliriumproducts/aumo/mysql"
	"github.com/gomodule/redigo/redis"
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

	a := auth.New(r, ustore, 60*60*24)

	t.Run("new_session", func(t *testing.T) {
		defer TidyRedis(r)
		user := createUser(t, ustore)

		sID, err := a.NewSession(user)
		require.Nil(t, err, "shouldn't return an error")
		require.NotEmpty(t, sID, "should return a session ID")

		uID, err := redis.Uint64(r.Do("GET", sID))
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, user.ID, uint(uID))
	})

	t.Run("get_user_from_session", func(t *testing.T) {
		defer TidyRedis(r)

		user := createUser(t, ustore)
		sID := createSession(t, r, user, 60*60*24)

		gotUser, err := a.Get(sID)
		require.Nil(t, err, "shouldn't return an error")
		require.Equal(t, user.ID, gotUser.ID)
	})

	t.Run("delete_session", func(t *testing.T) {
		defer TidyRedis(r)

		user := createUser(t, ustore)
		sID := createSession(t, r, user, 60*60*24)

		err = a.Del(sID)
		require.Nil(t, err, "shouldn't return an error")

		uID, err := redis.Uint64(r.Do("GET", sID))
		require.NotNil(t, err, "should return an error")
		require.Empty(t, uID, "shouldn't return a user ID")
	})
}
