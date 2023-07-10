package eng

import (
	"github.com/TheDevtop/exm/internal/fsio"
	"github.com/TheDevtop/exm/internal/logger"
	"github.com/TheDevtop/exm/shared"
)

// Update or create table
func UpCreate(table string, entry string, values []string) {
	var (
		tab shared.Table
		err error
	)

	if tab, err = fsio.ReadTable(table); err != nil {
		logger.ErrorLog.Printf("%s, using blank table", err)
		tab = make(shared.Table, 1)
	}
	tab[entry] = values
	if err = fsio.WriteTable(table, tab); err != nil {
		logger.ErrorLog.Println(err)
	}
}

func DeleteTable(table string) {
	if err := fsio.DeleteTable(table); err != nil {
		logger.ErrorLog.Println(err)
	}
}

func DeleteEntry(table string, entry string) {
	var (
		tab shared.Table
		err error
	)

	if tab, err = fsio.ReadTable(table); err != nil {
		logger.ErrorLog.Println(err)
		return
	}
	delete(tab, entry)
	if err = fsio.WriteTable(table, tab); err != nil {
		logger.ErrorLog.Println(err)
	}
}
