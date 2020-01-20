package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Error is a JSON type for error handling
type Error struct {
	Err string `json:"error"`
}

func (e Error) Error() string {
	return e.Err
}

var errMarshaling, _ = json.Marshal(Error{
	Err: "Failed to Marshal Error",
})

// JSONError is a convenience function for handling errors
func (rest *Rest) JSONError(w http.ResponseWriter, err error, statusCode int) {
	json, err := json.Marshal(Error{
		Err: err.Error(),
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

// Message is a struct for JSON responses in the format of
//    {
//        message: "User sucessfully deleted!"
//    }
type Message struct {
	Message string `json:"message"`
}

// JSON is a convenience function for writing to JSON
func (rest *Rest) JSON(w http.ResponseWriter, v interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := enc.Encode(v); err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(buf.Bytes())
}
