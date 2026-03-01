package handler

import (
	"net/http"
)

type PasswordHandler struct {
	Password string
}

func (h *PasswordHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "static/login.html")

	case http.MethodPost:

	}
}
