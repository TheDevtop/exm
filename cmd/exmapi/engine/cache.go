package engine

import (
	"regexp"
	"time"

	"github.com/TheDevtop/exm/lib"
)

var (
	reCache map[string]*regexp.Regexp
	config  lib.Config
)

func alloc() {
	reCache = make(map[string]*regexp.Regexp, config.CacheSize)
}

func clean() {
	for {
		if len(reCache) > config.CacheSize {
			alloc()
		}
		time.Sleep(time.Hour)
	}
}

func getRegex(restr string) (*regexp.Regexp, error) {
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

func Start(cfg lib.Config) {
	config = cfg
	alloc()
	go clean()
}
