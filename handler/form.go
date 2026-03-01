package handler

import (
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}
}
