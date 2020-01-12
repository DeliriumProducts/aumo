package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Error is a json type for error handling
type Error struct {
	Error string `json:"error"`
}

var errMarshaling, _ = json.Marshal(Error{
	Error: "Failed to Marshal Error",
})

// JSONError is a convenience function for handling errors
func JSONError(w http.ResponseWriter, err error, statusCode int) {
	json, err := json.Marshal(Error{
		Error: err.Error(),
	})

	// We fallback to a default error if we encountered one
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write(errMarshaling)
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write(json)
}

// JSON is a convenience function for writing to JSON
func JSON(w http.ResponseWriter, v interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := enc.Encode(v); err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(buf.Bytes())
}
