package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/TheDevtop/exm/api"
	"github.com/TheDevtop/exm/rec"
	"github.com/TheDevtop/exm/sti"
	"github.com/TheDevtop/go-probes"
)

func main() {
	var (
		pb  = probes.NewLogProbe("main.main", os.Stderr)
		err error
	)

	// Declare and parse flags
	flagDriver := flag.String("driver", "none", "Specify storage driver")
	flag.Parse()

	// Initialize the storage interface
	if err = sti.Setup(*flagDriver); err != nil {
		pb.Probe(err.Error())
		os.Exit(1)
	}

	// Initialize cache and API
	rec.Setup()
	api.Setup()

	// Get ready to service
	pb.Probe(fmt.Sprintf("Servicing on (%s)", api.ListenAddr))
	if err = http.ListenAndServe(api.ListenAddr, nil); err != nil {
		pb.Probe(err.Error())
	}
}
