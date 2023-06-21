package lib

import (
	"encoding/json"
	"os"
	"time"
)

type SourceEntry struct {
	Type    string
	Address string
	Timeout time.Duration
	Object  string
}

type Config struct {
	Directory     string
	CacheSize     int
	ServerAddress string
	Sources       []SourceEntry
}

func LoadConfig(path string) (Config, error) {
	var (
		cfg Config
		err error
		buf []byte
	)
	if buf, err = os.ReadFile(path); err != nil {
		return cfg, err
	}
	if err = json.Unmarshal(buf, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
