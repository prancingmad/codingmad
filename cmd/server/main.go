package main

import (
    "fmt"
    "log"
    "net/http"

	"codingmad/internal/config"
    "codingmad/internal/db"
)

func main() {
	cfg := config.Load()

    db.Connect(cfg.DBPath)
    defer db.Close()

    mux := http.NewServeMux()
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        fmt.Fprint(w, `{"status":"ok"}`)
    })

    fmt.Printf("Server running on http://localhost:%s\n", cfg.Port)
    if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
