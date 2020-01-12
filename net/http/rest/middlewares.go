package rest

import (
	"net/http"
)

// ParseForm is a middleware which calls r.ParseForm
func ParseForm(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			JSONError(w, err, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func (rest *Rest) WithAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Get the session from the store (cookie store)
// 		session, err := rest.store.Get(r, CookieStoreKey)
// 		if err != nil {
// 			http.Error(w, "User unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		// Retrieve our struct and type-assert it
// 		val := session.Values[UserSessionKey]
// 		user, ok := val.(aumo.User)
// 		if !ok {
// 			http.Error(w, "Bad Request", http.StatusBadRequest)
// 			return
// 		}

// 		// Update our user
// 		user, err = wb.Aumo.GetUserByID(user.ID)
// 		if err != nil {
// 			http.Error(w, "User not found", http.StatusNotFound)
// 			return
// 		}

// 		session.Values[UserSessionKey] = &user

// 		// Re-update the session (update cookie)
// 		err = session.Save(r, w)
// 		if err != nil {
// 			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		}

// 		next.ServeHTTP(w, r.WithContext(
// 			context.WithValue(r.Context(), UserContextKey, user),
// 		))

// 	})
// }
