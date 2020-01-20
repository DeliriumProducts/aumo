package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) userRegister(w http.ResponseWriter, r *http.Request) {
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

func (rest *Rest) userLogin(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email    string `form:"email" validate:"required,email" json:"email"`
		Password string `form:"password" validate:"required" json:"password"`
	}

	var um request
	if ok := rest.Form(w, r, &um); !ok {
		return
	}

	user, err := rest.userService.UserByEmail(um.Email, false)
	if err != nil {
		rest.JSONError(w, aumo.ErrUserNotFound, http.StatusNotFound)
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

func (rest *Rest) userGetCurrent(w http.ResponseWriter, r *http.Request) {
	user, err := auth.CurrentUser(r.Context())
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, user, http.StatusOK)
}

func (rest *Rest) userLogout(w http.ResponseWriter, r *http.Request) {
	sID, err := r.Cookie(auth.CookieKey)

	if err != nil {
		rest.JSONError(w, err, http.StatusUnauthorized)
		return
	}

	err = rest.auth.Del(sID.Value)

	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, Message{"User successfully logged out!"}, 200)
}
