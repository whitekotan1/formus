package handler

import (
	"net/http"
)

func LandingHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "static/index.html")

}
