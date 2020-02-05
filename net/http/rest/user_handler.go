package rest

import (
	"net/http"

	"github.com/deliriumproducts/aumo"
)

func (rest *Rest) userGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := rest.userService.Users()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, users, http.StatusOK)
}

func (rest *Rest) userGet(w http.ResponseWriter, r *http.Request) {
	uID := rest.Param(r, "id")

	user, err := rest.userService.User(uID, true)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, user, http.StatusOK)
}

func (rest *Rest) userEditRole(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Role aumo.Role `form:"role" validate:"required,oneof=Admin Customer" json:"role"`
	}

	uID := rest.Param(r, "id")

	var ur request
	if ok := rest.Form(w, r, &ur); !ok {
		return
	}

	err := rest.userService.EditRole(uID, ur.Role)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, Message{"User role successfully edited!"}, http.StatusOK)
}

func (rest *Rest) userAddPoints(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Points float64 `form:"points" validate:"required,numeric" json:"points"`
	}

	uID := rest.Param(r, "id")

	var ur request
	if ok := rest.Form(w, r, &ur); !ok {
		return
	}

	err := rest.userService.AddPoints(uID, ur.Points)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, Message{"User points successfully added!"}, http.StatusOK)
}

func (rest *Rest) userSubPoints(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Points float64 `form:"points" validate:"required,numeric" json:"points"`
	}

	uID := rest.Param(r, "id")

	var ur request
	if ok := rest.Form(w, r, &ur); !ok {
		return
	}

	err := rest.userService.SubPoints(uID, ur.Points)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, Message{"User points successfully subtracted!"}, http.StatusOK)
}

func (rest *Rest) userDelete(w http.ResponseWriter, r *http.Request) {
	uID := rest.Param(r, "id")

	err := rest.userService.Delete(uID)
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, Message{"User successfully deleted!"}, http.StatusOK)
}
