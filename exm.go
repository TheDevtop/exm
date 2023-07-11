package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TheDevtop/exm/internal/api"
	"github.com/TheDevtop/exm/shared"
)

func initLogger(logPtr *log.Logger) {
	logPtr.SetOutput(os.Stderr)
	logPtr.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	initLogger(log.Default())
	log.Println("Initialized logger ✅")

	if err := os.Chdir(os.Getenv(shared.Direnv)); err != nil {
		log.Fatalf("%s, please declare the %s variable ❌\n", err, shared.Direnv)
	}
	log.Printf("Changed directory to %s ✅\n", os.Getenv(shared.Direnv))

	api.MountRoutes(http.DefaultServeMux)
	log.Println("Mounted API routes ✅")

	log.Printf("Serving on address %s\n", shared.Port)
	http.ListenAndServe(shared.Port, nil)
}
