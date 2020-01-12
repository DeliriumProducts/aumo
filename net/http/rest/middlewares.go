package rest

import (
	"net/http"
)

const (
	xFrameOptions                = "X-Frame-Options"
	xFrameOptionsValue           = "DENY"
	xContentTypeOptions          = "X-Content-Type-Options"
	xContentTypeOptionsValue     = "nosniff"
	xssProtection                = "X-XSS-Protection"
	xssProtectionValue           = "1; mode=block"
	strictTransportSecurity      = "Strict-Transport-Security"                    // details https://blog.bracelab.com/achieving-perfect-ssl-labs-score-with-go + https://developer.mozilla.org/en-US/docs/Web/Security/HTTP_strict_transport_security
	strictTransportSecurityValue = "max-age=31536000; includeSubDomains; preload" // 31536000 = just shy of 12 months
	// also look at Content-Security-Policy in the future.
)

// Security Adds HTTP headers for XSS Protection and alike.
func Security(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(xFrameOptions, xFrameOptionsValue)
		w.Header().Add(xContentTypeOptions, xContentTypeOptionsValue)
		w.Header().Add(xssProtection, xssProtectionValue)
		w.Header().Add(strictTransportSecurity, strictTransportSecurityValue)

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
