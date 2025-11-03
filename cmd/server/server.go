package server

import (
	"fmt"
	"log"
	"net/http"

	"codingmad/internal/config"
	"codingmad/internal/db"
	"codingmad/internal/handlers"
)

func Start(cfg *config.Config) {
	// Connect SQLite database
	db.Connect(cfg.DBPath)
	defer db.Close()

	mux := http.NewServeMux()

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	// Register /notes routes
	handlers.RegisterNoteRoutes(mux)

	fmt.Printf("🚀 Server running on http://localhost:%s\n", cfg.Port)
	// This will block forever and keep the server running
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
