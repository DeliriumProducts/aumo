package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/deliriumproducts/aumo"
)

var (
	ErrBadTypeAssertion = errors.New("aumo: failed to assert type")
)

type RegisterForm struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=24"`
	Avatar   string `form:"avatar" validate:"required,url"`
}

func (rest *Rest) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var um RegisterForm
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := aumo.NewUser(um.Name, um.Email, um.Password, um.Avatar)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	err = rest.userService.Create(user)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	JSON(w, user, 200)
}

type LoginForm struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

func (rest *Rest) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var um LoginForm
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := rest.userService.UserByEmail(um.Email, false)
	if err != nil {
		JSONError(w, err, http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(um.Password) {
		JSONError(w, aumo.ErrInvalidPassword, http.StatusUnauthorized)
		return
	}

	sID, err := rest.auth.NewSession(user)
	if err != nil {
		JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.auth.SetCookieHeader(w, sID)
	JSON(w, user, 200)
}

func (rest *Rest) Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}
