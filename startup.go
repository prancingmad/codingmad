package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"codingmad/internal/config"
	"codingmad/internal/db"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Connect SQLite database
	db.Connect(cfg.DBPath)
	// defer db.Close() <- handled by server process; do not close here

	// Open browser after DB is ready
	url := "http://localhost:" + cfg.Port + "/health"
	openBrowser(url)

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	fmt.Printf("🚀 Server running on %s\n", url)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// openBrowser opens the URL in the default Windows browser (WSL compatible)
func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		if _, errWSL := os.Stat("/proc/version"); errWSL == nil {
			// WSL detected, use PowerShell to open Windows default browser
			err = exec.Command("powershell.exe", "Start-Process", url).Start()
		} else {
			// Linux fallback
			err = exec.Command("xdg-open", url).Start()
		}
	default:
		fmt.Println("Open browser manually:", url)
	}

	if err != nil {
		fmt.Println("Failed to open browser:", err)
	} else {
		fmt.Println("Opened browser to", url)
	}
}
