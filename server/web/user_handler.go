package web

import (
	"encoding/json"
	"net/http"
)

type UserRegisterForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (wb *Web) Register(w http.ResponseWriter, r *http.Request) {
	var urm UserRegisterForm
	json.Marshal()
}
