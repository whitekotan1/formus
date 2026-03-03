package handler

import (
	"formus/database"
	"html/template"
	"log"
	"net/http"
)

func AdminPanelHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
		return
	}

	forms := database.GetForms()

	template, err := template.ParseFiles("static/adminPanel.html")

	if err != nil {
		log.Println("error parsing template", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, forms)

	if err != nil {
		log.Println("rendering error", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
