package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	var err error

	DB, err = sql.Open("sqlite", "./formus.db")

	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	DB.Exec("PRAGMA foreign_keys = ON;")

	query := `
	CREATE TABLE IF NOT EXISTS forms (
    id INTEGER  PRIMARY KEY AUTOINCREMENT, 
    name TEXT NOT NULL
);


	CREATE TABLE IF NOT EXISTS requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    form_id INTEGER NOT NULL,
    data STRING NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (form_id) REFERENCES forms(id)

)
	`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("can't create tables", err)
	}

	log.Println("succesfully initialize db")

}
