package runners

import (
	"io"
	"log"
	"net"
	"time"

	"github.com/TheDevtop/exm/lib"
)

const TcpType = "tcp"

var TcpRunner lib.Runner = func(se lib.SourceEntry, wf func(string, io.ReadCloser)) {
	if se.Timeout < 1 {
		se.Timeout = time.Minute
	}
	for {
		if sd, err := net.Dial(TcpType, se.Address); err != nil {
			log.Printf(errDialing, err, se.Address)
			sd.Close()
		} else {
			wf(se.Object, sd)
		}
		time.Sleep(se.Timeout * time.Second)
	}
}
