package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserRegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (wb *Web) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	var urm UserRegisterForm
	json.Unmarshal(body, &urm)

	user, err := wb.CreateUser(urm.Name, urm.Email, urm.Password)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	res, err := json.Marshal(user)

	w.Write(res)
}
