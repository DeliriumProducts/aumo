package web

import (
	"encoding/json"
	"net/http"

	"github.com/fr3fou/aumo/server/aumo"
)

type UserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (wb *Web) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := wb.CreateUser(um.Name, um.Email, um.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm

	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := wb.store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := wb.GetUserByEmail(um.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(um.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
	}

	session.Values["user"] = &user

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wb *Web) SecretHandler(w http.ResponseWriter, r *http.Request) {
	session, err := wb.store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve our struct and type-assert it
	val := session.Values["user"]
	user, ok := val.(aumo.User)
	if !ok {
		http.Error(w, "User unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
