package cache

import "github.com/TheDevtop/exm/pkg/data"

var LocalCache data.Table

func init() {
	LocalCache = make(data.Table, data.CacheSize)
}
