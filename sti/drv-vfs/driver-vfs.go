package drvvfs

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TheDevtop/go-probes"
)

const DriverName = "vfs"

var vfsDir = os.Getenv("VFSDIR")

func Stream(path string) (*bufio.Scanner, error) {
	var (
		pb  = probes.NewLogProbe("drvvfs.Stream", os.Stderr)
		fd  *os.File
		err error
	)

	if fd, err = os.Open(path); err != nil {
		pb.Probe(err.Error())
		return nil, err
	}
	return bufio.NewScanner(fd), nil
}

func List() ([]string, error) {
	var (
		pb      = probes.NewLogProbe("drvvfs.List", os.Stderr)
		list    []string
		entries []os.DirEntry
		err     error
	)

	if entries, err = os.ReadDir(vfsDir); err != nil {
		pb.Probe(err.Error())
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			list = append(list, entry.Name())
		}
	}
	return list, nil
}

func Setup() error {
	pb := probes.NewLogProbe("drvvfs.Setup", os.Stderr)
	if err := os.Chdir(vfsDir); err != nil {
		pb.Probe(err.Error())
		return err
	}
	pb.Probe(fmt.Sprintf("Successful chdir (%s)", vfsDir))
	return nil
}
