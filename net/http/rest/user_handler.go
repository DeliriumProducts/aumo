package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/deliriumproducts/aumo"
)

var (
	ErrBadTypeAssertion = errors.New("aumo: failed to assert type")
)

type UserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

func (rest *Rest) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := aumo.NewUser(um.Name, um.Email, um.Password, um.Avatar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rest.UserService.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (rest *Rest) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm

	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if !user.ValidatePassword(um.Password) {
	// 	http.Error(w, "Invalid password", http.StatusUnauthorized)
	// }
}
