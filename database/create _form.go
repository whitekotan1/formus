package database

import (
	_ "modernc.org/sqlite"
)

func CreateForm(name string) error {
	query := `INSERT INTO forms (name) VALUES(?);   `

	_, err := DB.Exec(query, name)
	return err
}
