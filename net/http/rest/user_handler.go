package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/deliriumproducts/aumo"
)

var (
	ErrBadTypeAssertion = errors.New("aumo: failed to assert type")
)

type UserForm struct {
	Username string `form:"username"`
	Email    string `form:"email"`
	Avatar   string `form:"avatar"`
	Password string `form:"password"`
}

func (rest *Rest) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if err := rest.decoder.Decode(&um, r.Form); err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := aumo.NewUser(um.Username, um.Email, um.Password, um.Avatar)
	if err != nil {
		JSONError(w, err, http.StatusBadRequest)
		return
	}
	if err := rest.validator.Struct(user); err != nil {
		spew.Dump(err)

		JSONError(w, err, http.StatusBadRequest)
		return
	}

	err = rest.userService.Create(user)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
	}

	JSON(w, user, 200)
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
