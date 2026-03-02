package database

import (
	_ "modernc.org/sqlite"
)

type Requests struct {
	ID        int
	FormID    int
	Data      string
	CreatedAt string
}

func SaveRequest(formID string, jsonData []byte) error {
	query := `INSERT INTO requests (form_id, data) VALUES(?,?)`

	_, err := DB.Exec(query, formID, string(jsonData))
	return err
}
