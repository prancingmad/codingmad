package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect(dbPath string) *sql.DB {
	var err error
	DB, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("SQLite Database connected successfully")
	return DB
}

func Close() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}