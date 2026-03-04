package database

import (
	"log"

	_ "modernc.org/sqlite"
)

type Request struct {
	ID          int
	FormID      int
	Data        string
	CreatedTime string
}

func GetRequests(formID string) []Request {
	var requests []Request

	dbRequest := (`
	SELECT id, form_id, data, created_at
	FROM requests
	WHERE form_id = ?
	 ORDER BY id DESC
	 `)

	query, err := DB.Query(dbRequest, formID)

	if err != nil {
		log.Println("can,t read from db", err)
		return requests
	}

	defer query.Close()

	for query.Next() {
		request := Request{}

		err := query.Scan(&request.ID, &request.FormID, &request.Data, &request.CreatedTime)

		if err != nil {
			log.Println("can't read row from db", err)
			continue
		}
		requests = append(requests, request)

	}
	return requests
}
