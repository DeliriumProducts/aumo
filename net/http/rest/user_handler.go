package rest

import "net/http"

func (rest *Rest) userGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := rest.userService.Users()
	if err != nil {
		rest.JSONError(w, err, http.StatusNotFound)
		return
	}

	rest.JSON(w, users, http.StatusOK)
}
