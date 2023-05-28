package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheDevtop/exm/cmd/exmsrv/runners"
	"github.com/TheDevtop/exm/lib"
)

const (
	infoConfigured = "Info: Configured object (%s) with runner (%s)\n"
)

func writeStream(object string, stream io.ReadCloser) {
	var (
		err error
		fd  *os.File
	)

	if fd, err = os.Create(object); err != nil {
		log.Printf(lib.ErrObject, err, object)
		return
	}
	if _, err := io.Copy(fd, stream); err != nil {
		log.Printf(lib.ErrObject, err, object)
		return
	}

	fd.Close()
	log.Printf("Info: Written object (%s) to store\n", object)
}

func main() {
	var (
		cfg   lib.Config
		err   error
		sigch = make(chan os.Signal, 1)
	)

	if len(os.Args) < 2 {
		log.Fatalln(lib.ErrNoConfig)
	}

	if cfg, err = lib.LoadConfig(os.Args[1]); err != nil {
		log.Fatalf(lib.ErrCantStart, err)
	}
	if err = os.Chdir(cfg.Directory); err != nil {
		log.Fatalf(lib.ErrCantStart, err)
	}

	for _, entry := range cfg.Sources {
		go func(centry lib.SourceEntry) {
			switch centry.Type {
			case runners.HttpType:
				log.Printf(infoConfigured, centry.Object, runners.HttpType)
				go runners.HttpRunner(centry, writeStream)
			case runners.TcpType:
				log.Printf(infoConfigured, centry.Object, runners.TcpType)
				go runners.TcpRunner(centry, writeStream)
			case runners.UdpType:
				log.Printf(infoConfigured, centry.Object, runners.UdpType)
				go runners.UdpRunner(centry, writeStream)
			case runners.LocalType:
				log.Printf(infoConfigured, centry.Object, runners.LocalType)
			default:
				log.Printf("Error: Object (%s) has type (%s), unknown runner\n", centry.Object, centry.Type)
			}
		}(entry)
	}

	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	log.Printf(lib.InfoSignal, <-sigch)
	os.Exit(0)
}
