package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) registerHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name     string `form:"name" validate:"required"`
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required,min=6,max=24"`
		Avatar   string `form:"avatar" validate:"required,url"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := aumo.NewUser(um.Name, um.Email, um.Password, um.Avatar)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	err = rest.userService.Create(user)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, user, http.StatusOK)
}

func (rest *Rest) loginHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email    string `form:"email" validate:"required,email"`
		Password string `form:"password" validate:"required"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := rest.userService.UserByEmail(um.Email, false)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(um.Password) {
		rest.JSONError(w, aumo.ErrInvalidPassword, http.StatusUnauthorized)
		return
	}

	sID, err := rest.auth.NewSession(user)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.auth.SetCookieHeader(w, sID)
	rest.JSON(w, user, http.StatusOK)
}

func (rest *Rest) meHandler(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUserFromContext(r.Context())
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, user, http.StatusOK)
}
