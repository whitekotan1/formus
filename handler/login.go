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

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		inputedPassword := r.FormValue("password")

		if inputedPassword != h.Password {
			http.Error(w, "invalid password", http.StatusUnauthorized)
			return
		}

		cookie := &http.Cookie{
			Name:     "formus_auth",
			Value:    "authorized_true",
			Path:     "/form",
			MaxAge:   3600 * 200,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/adminPanel", http.StatusFound)

	default:
		http.Error(w, "method not aloowed", http.StatusMethodNotAllowed)

	}
}
