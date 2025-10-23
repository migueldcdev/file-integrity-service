package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectToDB(path string) error {
	var err error
	DB, err = sql.Open("sqlite3", path)
	if err != nil {
		return fmt.Errorf("Failed to open db %v\n", err)
	}
	return nil
}

func InitDB() error {
	var err error
	createTable := `
        CREATE TABLE IF NOT EXISTS hash (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            path TEXT UNIQUE NOT NULL,
            hash TEXT NOT NULL,
            size INTEGER NOT NULL,
            created_at DATETIME NOT NULL,
            last_checked DATETIME NOT NULL
        )
       `
	_, err = DB.Exec(createTable)

	if err != nil {
		return fmt.Errorf("Error while creating table: %v\n", err)
	}

	return nil

}
