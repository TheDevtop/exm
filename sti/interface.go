package sti

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	drvminio "github.com/TheDevtop/exm/sti/drv-minio"
	drvmockup "github.com/TheDevtop/exm/sti/drv-mockup"
	drvvfs "github.com/TheDevtop/exm/sti/drv-vfs"
	"github.com/TheDevtop/go-probes"
)

// Functions that provide the storage interface
var Stream func(string) (*bufio.Scanner, error)
var List func() ([]string, error)

func Setup(driver string) error {
	var (
		pb        = probes.NewLogProbe("sti.Setup", os.Stderr)
		err error = nil
	)

	switch driver {
	case drvminio.DriverName:
		err = drvminio.Setup()
		Stream = drvminio.Stream
		List = drvminio.List
		pb.Probe(fmt.Sprintf("Storage driver (%s)", drvminio.DriverName))
	case drvvfs.DriverName:
		err = drvvfs.Setup()
		Stream = drvvfs.Stream
		List = drvvfs.List
		pb.Probe(fmt.Sprintf("Storage driver (%s)", drvvfs.DriverName))
	case drvmockup.DriverName:
		err = drvmockup.Setup()
		Stream = drvmockup.Stream
		List = drvmockup.List
		pb.Probe(fmt.Sprintf("Storage driver (%s)", drvmockup.DriverName))
	default:
		err = errors.New("driver not found error")
		pb.Probe(err.Error())
	}
	return err
}
