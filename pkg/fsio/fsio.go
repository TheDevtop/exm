package fsio

import (
	"os"

	"github.com/TheDevtop/exm/pkg/data"
	toml "github.com/pelletier/go-toml/v2"
)

// Write table to filesystem
func WriteTable(path string, table data.Table) error {
	if buf, err := toml.Marshal(table); err != nil {
		return err
	} else if err := os.WriteFile(path, buf, data.DefaultPerm); err != nil {
		return err
	}
	return nil
}

// Read table from filesystem
func ReadTable(path string) (data.Table, error) {
	var tab data.Table
	if buf, err := os.ReadFile(path); err != nil {
		return nil, err
	} else if toml.Unmarshal(buf, tab); err != nil {
		return nil, err
	} else {
		return tab, nil
	}
}
