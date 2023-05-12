package drvmockup

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TheDevtop/go-probes"
)

const DriverName string = "mockup"

var mockStore = map[string]string{
	"foobar": "This is the content of foobar",
	"foobaz": "This is the content of foobaz",
}

func Setup() error {
	probes.NewLogProbe("drvmockup.Setup", os.Stderr).Probe("Mockup driver initialized")
	return nil
}

func Stream(id string) (*bufio.Scanner, error) {
	pb := probes.NewLogProbe("drvmockup.Stream", os.Stderr)
	if str, ok := mockStore[id]; !ok {
		err := fmt.Errorf("object %s not found", id)
		pb.Probe(err.Error())
		return nil, err
	} else {
		rd := strings.NewReader(str)
		return bufio.NewScanner(rd), nil
	}
}

func List() ([]string, error) {
	var list []string
	for k := range mockStore {
		list = append(list, k)
	}
	return list, nil
}
