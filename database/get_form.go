package database

import (
	"log"

	_ "modernc.org/sqlite"
)

type Form struct {
	ID   int
	Name string
}

func GetForms() []Form {

	var forms []Form

	query, err := DB.Query("SELECT id, name FROM forms ORDER BY id DESC")

	if err != nil {
		log.Println("can't read from db", err)
		return forms
	}

	defer query.Close()

	for query.Next() {
		form := Form{}

		err := query.Scan(&form.ID, &form.Name)

		if err != nil {
			log.Println("can't get row from db", err)
			continue
		}

		forms = append(forms, form)
	}
	return forms
}
