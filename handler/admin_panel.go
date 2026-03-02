package handler

import (
	"net/http"
)

func AdminPanelHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "static/adminPanel.html")
}
