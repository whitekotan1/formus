package handler

import (
	"formus/database"
	"net/http"
)

func CreateFormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "can't parse form", http.StatusBadRequest)
		return
	}

	formName := r.FormValue("name")

	if formName == "" {
		http.Error(w, "form name can't be empty", http.StatusBadRequest)
		return
	}

	err = database.CreateForm(formName)

	if err != nil {
		http.Error(w, "can't write form name to db", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
