package engine

import (
	"bufio"
	"errors"
	"io/fs"
	"os"
	"strings"
)

const (
	charSlash       = "/"
	errContainSlash = "object name contains \"/\""
)

func getObject(object string) (*bufio.Scanner, error) {
	if strings.ContainsAny(object, charSlash) {
		return nil, errors.New(errContainSlash)
	}
	if fd, err := os.Open(object); err != nil {
		return nil, err
	} else {
		return bufio.NewScanner(fd), nil
	}
}

func getMetadata(object string) (fs.FileInfo, error) {
	if strings.ContainsAny(object, charSlash) {
		return nil, errors.New(errContainSlash)
	}
	if fi, err := os.Stat(object); err != nil {
		return nil, err
	} else {
		return fi, nil
	}
}
