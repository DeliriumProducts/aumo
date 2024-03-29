package rest

import (
	"errors"
	"net/http"
	"time"

	"github.com/deliriumproducts/aumo"
	"github.com/deliriumproducts/aumo/auth"
)

func (rest *Rest) userRegister(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name     string `form:"name" validate:"required" json:"name"`
		Email    string `form:"email" validate:"required,email" json:"email"`
		Password string `form:"password" validate:"required,min=6,max=24" json:"password"`
		Avatar   string `form:"avatar" validate:"required,url" json:"avatar"`
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
	switch {
	case err == nil:
		break
	case errors.Is(err, aumo.ErrDuplicateEmail):
		rest.JSONError(w, err, http.StatusUnprocessableEntity)
		return
	default:
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = rest.verifier.Send(user.Email, user.ID.String(), "Aumo Confirmation Email", "This is an email for confirming your Aumo account", rest.backendURL+"/confirm-email", time.Hour*24)
	if err != nil {
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	rest.JSON(w, Message{"Email confirmation sent!"}, http.StatusOK)
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

	if !user.IsVerified {
		rest.JSONError(w, aumo.ErrNotVerified, http.StatusUnauthorized)
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

func (rest *Rest) userConfirmEmail(w http.ResponseWriter, r *http.Request) {
	token := rest.Param(r, "token")
	if token == "" {
		rest.JSONError(w, Error{"Token not provided"}, http.StatusBadRequest)
		return
	}

	userID, err := rest.verifier.Verify(token)
	if err != nil {
		rest.JSONError(w, Error{"Token not found"}, http.StatusNotFound)
		return
	}

	err = rest.userService.Verify(userID)
	switch {
	case err == nil:
		break
	case errors.Is(err, aumo.ErrUserNotFound):
		rest.JSONError(w, err, http.StatusNotFound)
		return
	default:
		rest.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "https://aumo.deliprods.tech", 301)
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

	rest.auth.ExpireCookieHeader(w)
	rest.JSON(w, Message{"User successfully logged out!"}, http.StatusOK)
}
