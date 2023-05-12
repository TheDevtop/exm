package rec

import (
	"os"
	"regexp"
	"time"

	"github.com/TheDevtop/go-probes"
)

const (
	cacheMaxSize = 4  // Maximum size of cache
	cacheTimeout = 10 // Minutes to wait for timeout
)

var reCache map[string]*regexp.Regexp

// Reallocates cache
func clean() {
	pb := probes.NewLogProbe("rec.clean", os.Stderr)
	for {
		time.Sleep(cacheTimeout * time.Minute)
		if len(reCache) > cacheMaxSize {
			reCache = make(map[string]*regexp.Regexp, cacheMaxSize)
			pb.Probe("Cleaned the regex cache!")
		}
	}
}

// Allocates cache, starts cleanup function if selected
func Setup(autoclean bool) {
	pb := probes.NewLogProbe("rec.Setup", os.Stderr)
	reCache = make(map[string]*regexp.Regexp, cacheMaxSize)

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
