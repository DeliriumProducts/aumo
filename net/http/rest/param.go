package rest

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Param is method to automatically get the param
func (rest *Rest) Param(r *http.Request, param string) string {
	return chi.URLParam(r, param)
}

// ParamNumber is method to automatically get the param
// and convert it to uint
func (rest *Rest) ParamNumber(w http.ResponseWriter, r *http.Request, param string) uint {

	p := chi.URLParam(r, param)

	n, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		rest.JSONError(w, err, http.StatusBadRequest)
	}

	return uint(n)
}
