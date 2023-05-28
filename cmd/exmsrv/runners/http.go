package runners

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/TheDevtop/exm/lib"
)

const HttpType = "http"

var HttpRunner lib.Runner = func(se lib.SourceEntry, wf func(string, io.ReadCloser)) {
	if se.Timeout < 1 {
		se.Timeout = time.Minute
	}
	for {
		if resp, err := http.DefaultClient.Get(se.Address); err != nil {
			log.Printf(errDialing, err, se.Address)
		} else {
			wf(se.Object, resp.Body)
		}
		time.Sleep(se.Timeout * time.Second)
	}
}
