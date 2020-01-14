package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/net/http/rest/auth"
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

func (rest *Rest) WithAuth(roles ...aumo.Role) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := rest.auth.GetFromRequest(r)
			if err != nil {
				JSONError(w, Error{"User unauthorized"}, http.StatusUnauthorized)
				return
			}

			if len(roles) > 0 || user.Role != aumo.Admin {
				isAuthorized := false
				for _, role := range roles {
					if user.Role == role {
						isAuthorized = true
						break
					}
				}

				if !isAuthorized {
					JSONError(w, Error{"User unauthorized"}, http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r.WithContext(
				auth.SetUserToContext(r.Context(), *user),
			))
		})
	}
}
