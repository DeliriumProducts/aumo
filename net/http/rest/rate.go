package rest

import (
	"net/http"
)

func (rest *Rest) onRateLimit(w http.ResponseWriter, r *http.Request) {
	rest.JSONError(w, Error{"You have reached maximum request limit."}, http.StatusTooManyRequests)
}
