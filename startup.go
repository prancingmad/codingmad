package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"codingmad/cmd/server"
	"codingmad/internal/config"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Optional: do other startup tasks here
	// e.g., start Docker, RabbitMQ, etc. in the future

	// Open browser after everything is ready
	url := "http://localhost:" + cfg.Port + "/health"
	openBrowser(url)

	// Start the full server (blocks forever)
	server.Start(cfg)
}

// openBrowser opens the URL in the default browser
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
