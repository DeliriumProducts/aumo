package web

import (
	"encoding/json"
	"net/http"
)

type UserRegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (wb *Web) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var urm UserRegisterForm
	if err := json.NewDecoder(r.Body).Decode(&urm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := wb.CreateUser(urm.Name, urm.Email, urm.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (wb *Web) LoginHandler(w http.ResponseWriter, r *http.Response) {

}
