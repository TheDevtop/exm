package sti

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	drvminio "github.com/TheDevtop/exm/sti/drv-minio"
	drvvfs "github.com/TheDevtop/exm/sti/drv-vfs"
	"github.com/TheDevtop/go-probes"
)

// Return a new data streamer
var Stream func(string) (*bufio.Scanner, error)

func Setup(driver string) error {
	var (
		pb        = probes.NewLogProbe("sti.Setup", os.Stderr)
		err error = nil
	)

	switch driver {
	case drvminio.DriverName:
		err = drvminio.Setup()
		Stream = drvminio.Stream
		pb.Probe(fmt.Sprintf("Storage driver (%s)", drvminio.DriverName))
	case drvvfs.DriverName:
		err = drvvfs.Setup()
		Stream = drvvfs.Stream
		pb.Probe(fmt.Sprintf("Storage driver (%s)", drvvfs.DriverName))
	default:
		err = errors.New("driver not found error")
		pb.Probe(err.Error())
	}
	return err
}
