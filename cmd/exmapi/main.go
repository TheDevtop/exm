package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheDevtop/exm/cmd/exmapi/engine"
	"github.com/TheDevtop/exm/lib"
)

/*
Program entrypoint, load the configuration, start the engine;
mount endpoints, start the server.
Finally, listen to stop signals.
*/

func main() {
	var (
		err    error
		sigch  = make(chan os.Signal, 1)
		config lib.Config
	)

	if len(os.Args) < 2 {
		log.Fatalln(lib.ErrNoConfig)
	}

	if config, err = lib.LoadConfig(os.Args[1]); err != nil {
		log.Fatalf(lib.ErrCantStart, err)
	}
	if err = os.Chdir(config.Directory); err != nil {
		log.Fatalf(lib.ErrCantStart, err)
	}

	engine.Start(config)

	http.HandleFunc(urlSearch_object, apiSearch_object)
	http.HandleFunc(urlSearch_global, apiSearch_global)
	http.HandleFunc(urlIndex_object, apiIndex_object)
	http.HandleFunc(urlIndex_global, apiIndex_global)
	http.HandleFunc(urlMeta_object, apiMeta_object)
	http.HandleFunc(urlPing, apiPing)

	go http.ListenAndServe(config.ServerAddress, nil)

	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	log.Printf(lib.InfoSignal, <-sigch)
	os.Exit(0)
}
