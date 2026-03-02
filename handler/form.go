package handler

import (
	"encoding/json"
	"fmt"
	"formus/database"
	"net/http"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "only post method allowed", http.StatusMethodNotAllowed)
		return
	}

	formID := r.URL.Query().Get("form_id")

	if formID == "" {
		http.Error(w, "invalid form_id", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(r.PostForm)

	if err != nil {
		http.Error(w, "can't encode data", http.StatusInternalServerError)
		return
	}

	err = database.SaveRequest(formID, jsonData)

	if err != nil {
		fmt.Println("can't save form", err)
		return
	}
	fmt.Println("form was saved")

	w.WriteHeader(http.StatusOK)

}
