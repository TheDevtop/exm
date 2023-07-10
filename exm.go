package main

import (
	"net/http"
	"os"

	"github.com/TheDevtop/exm/internal/api"
	"github.com/TheDevtop/exm/internal/logger"
	"github.com/TheDevtop/exm/shared"
)

func main() {
	logger.InfoLog.Println("✨ Welcome to EXM! ✨")

	if err := os.Chdir(os.Getenv(shared.Direnv)); err != nil {
		logger.ErrorLog.Fatalf("%s, please declare the %s variable!\n", err, shared.Direnv)
	}
	logger.InfoLog.Printf("Changed directory to %s\n", os.Getenv(shared.Direnv))

	api.MountRoutes(http.DefaultServeMux)
	logger.InfoLog.Println("Mounted API routes")

	logger.InfoLog.Printf("Serving on address %s\n", shared.Port)
	http.ListenAndServe(shared.Port, nil)
}
