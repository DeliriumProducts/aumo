package web

import (
	"context"
	"net/http"

	"github.com/fr3fou/aumo/server/aumo"
)

// ContentTypeJSON is a middleware for setting the Content-Type header to application/json
func ContentTypeJSON(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func (wb *Web) WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session from the store (cookie store)
		session, err := wb.store.Get(r, CookieStoreKey)
		if err != nil {
			http.Error(w, "User unauthorized", http.StatusUnauthorized)
			return
		}

		// Retrieve our struct and type-assert it
		val := session.Values[UserSessionKey]
		user, ok := val.(aumo.User)
		if !ok {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Update our user
		user, err = wb.Aumo.GetUserByID(user.ID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		session.Values[UserSessionKey] = &user

		// Re-update the session (update cookie)
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		next.ServeHTTP(w, r.WithContext(
			context.WithValue(r.Context(), UserContextKey, user),
		))

	})
}

func GetUserFromContext(ctx context.Context) (aumo.User, error) {
	if user, ok := ctx.Value(UserContextKey).(aumo.User); ok {
		return user, nil
	}

	return aumo.User{}, ErrBadTypeAssertion
}
