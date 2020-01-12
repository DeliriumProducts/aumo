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
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=24"`
	Avatar   string `form:"avatar" validate:"required,url"`
}

func (rest *Rest) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := aumo.NewUser(um.Name, um.Email, um.Password, um.Avatar)
	err = rest.userService.Create(user)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
	}

	JSON(w, user, 200)
}

func (rest *Rest) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var um UserForm
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&um); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if !user.ValidatePassword(um.Password) {
	// 	http.Error(w, "Invalid password", http.StatusUnauthorized)
	// }
}
