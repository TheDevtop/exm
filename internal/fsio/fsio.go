package fsio

import (
	"fmt"
	"os"

	"github.com/TheDevtop/exm/shared"
	toml "github.com/pelletier/go-toml/v2"
)

const defaultPerm = 0644
const defaultExt = "%s.db"

// Delete table
func DeleteTable(name string) error {
	path := fmt.Sprintf(defaultExt, name)
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

// Write table to filesystem
func WriteTable(name string, table shared.Table) error {
	path := fmt.Sprintf(defaultExt, name)
	if buf, err := toml.Marshal(table); err != nil {
		return err
	} else if err := os.WriteFile(path, buf, defaultPerm); err != nil {
		return err
	}
	return nil
}

// Read table from filesystem
func ReadTable(name string) (shared.Table, error) {
	path := fmt.Sprintf(defaultExt, name)
	var tab shared.Table
	if buf, err := os.ReadFile(path); err != nil {
		return nil, err
	} else if toml.Unmarshal(buf, tab); err != nil {
		return nil, err
	} else {
		return tab, nil
	}
}
