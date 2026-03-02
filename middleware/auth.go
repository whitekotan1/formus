package middleware

import (
	"net/http"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("formus_auth")

		if err != nil || cookie.Value != "authorized_true" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return

		}

		next(w, r)
	}

}
