package lib

import (
	"encoding/json"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := Config{
		Directory:     "/foo/bar",
		CacheSize:     16,
		ServerAddress: ":1800",
	}
	cfg.Sources = []SourceEntry{
		{Type: "http/tcp/udp/local", Address: "", Timeout: 0, Object: "lorem.html"},
	}

	if buf, err := json.Marshal(&cfg); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%s\n", string(buf))
	}
}
