package rec

import (
	"os"
	"regexp"
	"time"

	"github.com/TheDevtop/go-probes"
)

const cacheTimeout = 15 // Minutes to wait for timeout

var (
	reCache  map[string]*regexp.Regexp
	cachSize int
)

// Reallocates cache
func clean() {
	pb := probes.NewLogProbe("rec.clean", os.Stderr)
	for {
		time.Sleep(cacheTimeout * time.Minute)
		if len(reCache) > cachSize {
			reCache = make(map[string]*regexp.Regexp, cachSize)
			pb.Probe("Cleaned the regex cache!")
		}
	}
}

// Allocates cache, starts autoclean if true
func Setup(size int, autoclean bool) {
	pb := probes.NewLogProbe("rec.Setup", os.Stderr)
	cachSize = size
	reCache = make(map[string]*regexp.Regexp, cachSize)

	// Is autoclean enabled
	if autoclean {
		go clean()
		return
	}
	pb.Probe("Warning, running without autoclean")
}

// Get regex from cache or compile on the spot
func Receive(restr string) (*regexp.Regexp, error) {
	if re, ok := reCache[restr]; ok {
		return re, nil
	}
	if re, err := regexp.Compile(restr); err != nil {
		return nil, err
	} else {
		reCache[restr] = re
		return re, nil
	}
}
