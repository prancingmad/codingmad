package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect(dbPath string) *sql.DB {
	if err := os.MkdirAll("./data", os.ModePerm); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("SQLite Database connected successfully")

	CreateTables(DB)

	return DB
}

func CreateTables(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS notes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		language TEXT
	);
	`

	if _, err := db.Exec(query); err != nil {
		log.Fatalf("Failed to create notes table: %v", err)
	}

	log.Println("Notes table ensured")
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
