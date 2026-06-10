package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitDB(path string) (*sql.DB, error) {

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := initDatabase(db); err != nil {
		return nil, err
	}
	return db, nil

}

func initDatabase(db *sql.DB) error {
	createTableSQ := `CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(40) NOT NULL,
    comment TEXT NOT NULL,
    repeat VARCHAR(128) NOT NULL
	);`

	if _, err := db.Exec(createTableSQ); err != nil {
		return err
	}

	createIndexSQL := `CREATE INDEX IF NOT EXISTS idx_scheduler_date ON scheduler(date);`

	_, err := db.Exec(createIndexSQL)
	return err
}
