package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents any error that may have occurred during validation
type ValidationError struct {
	Errors []string `json:"errors"`
}

// Form parses, validates, writes any errors that may have occurred during the process
// and returns if it succeeded or not
func (rest *Rest) Form(w http.ResponseWriter, r *http.Request, form interface{}) bool {
	switch contentType := r.Header.Get("Content-Type"); {
	case strings.HasPrefix(contentType, "application/json"):
		if err := json.NewDecoder(r.Body).Decode(form); err != nil {
			rest.JSONError(w, err, http.StatusBadRequest)
			return false
		}
	case
		strings.HasPrefix(contentType, "application/x-www-form-urlencoded"),
		strings.HasPrefix(contentType, "multipart/form-data"):
		if err := r.ParseForm(); err != nil {
			rest.JSONError(w, err, http.StatusBadRequest)
			return false
		}

		if err := rest.decoder.Decode(form, r.Form); err != nil {
			rest.JSONError(w, err, http.StatusBadRequest)
			return false
		}
	default:
		rest.JSONError(w, Error{"Bad Request"}, http.StatusBadRequest)
		return false
	}

	if err := rest.validator.Struct(form); err != nil {
		jsonErrs := ValidationError{
			Errors: []string{},
		}

		for _, e := range err.(validator.ValidationErrors) {
			jsonErrs.Errors = append(jsonErrs.Errors, e.Translate(rest.translator))
		}

		rest.JSON(w, jsonErrs, http.StatusBadRequest)
		return false
	}

	return true
}
