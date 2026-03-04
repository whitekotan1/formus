package handler

import (
	"encoding/json"
	"formus/database"
	"html/template"
	"log"
	"net/http"
)

type RequestView struct {
	ID          int
	CreatedTime string
	Fields      map[string]interface{}
}

func GetRequestsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only get method allowed", http.StatusMethodNotAllowed)
		return
	}

	formID := r.URL.Query().Get("form_id")
	if formID == "" {
		http.Error(w, "empty form_id", http.StatusBadRequest)
		return
	}

	rawRequests := database.GetRequests(formID)

	var viewRequests []RequestView

	for _, req := range rawRequests {
		var fields map[string]interface{}

		err := json.Unmarshal([]byte(req.Data), &fields)
		if err != nil {
			log.Println("json parse error:", err)
			continue
		}

		viewRequests = append(viewRequests, RequestView{
			ID:          req.ID,
			CreatedTime: req.CreatedTime,
			Fields:      fields,
		})
	}

	tmpl, err := template.ParseFiles("static/requestsDashboard.html")
	if err != nil {
		log.Println("can't parse requests template", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, viewRequests)
	if err != nil {
		log.Println("rendering error", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}
