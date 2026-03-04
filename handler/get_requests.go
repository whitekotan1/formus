package handler

import (
	"formus/database"
	"html/template"
	"log"
	"net/http"
)

func GetRequestsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowd", http.StatusMethodNotAllowed)
		return
	}

	formID := r.URL.Query().Get("form_id")

	if formID == "" {
		http.Error(w, "empty form_id", http.StatusBadRequest)
		return
	}

	requests := database.GetRequests(formID)

	template, err := template.ParseFiles("static/requestsDashboard.html")

	if err != nil {
		log.Println("can't parse requests template", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, requests)

	if err != nil {
		log.Println("rendering error", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

}
