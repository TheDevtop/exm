package main

import (
	"net/http"
	"os"

	"github.com/TheDevtop/exm/api"
	"github.com/TheDevtop/exm/conio"
)

func main() {
	const (
		fprobe = "main.main"
		port   = ":1800"
	)

	if err := conio.Setup(
		os.Getenv("S3HOST"),
		os.Getenv("S3USER"),
		os.Getenv("S3SECRET"),
		os.Getenv("S3BUCKET")); err != nil {
		conio.Probeln(fprobe, err.Error())
		os.Exit(1)
	}

	api.Setup()
	conio.Probeln(fprobe, ("Listening on " + port))
	http.ListenAndServe(port, nil)
	os.Exit(0)
}
