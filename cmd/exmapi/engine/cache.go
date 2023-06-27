package engine

import (
	"regexp"

	"github.com/TheDevtop/exm/lib"
	"github.com/zyedidia/generic/cache"
)

var (
	reCache *cache.Cache[string, *regexp.Regexp]
	config  lib.Config
)

func getRegex(restr string) (*regexp.Regexp, error) {
	if re, ok := reCache.Get(restr); ok {
		return re, nil
	}
	if re, err := regexp.Compile(restr); err != nil {
		return nil, err
	} else {
		reCache.Put(restr, re)
		return re, nil
	}
}

func Start(cfg lib.Config) {
	config = cfg
	reCache = cache.New[string, *regexp.Regexp](cfg.CacheSize)
}
