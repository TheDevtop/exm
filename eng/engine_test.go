package eng

import (
	"bufio"
	"regexp"
	"testing"

	"github.com/TheDevtop/exm/rec"
	drvmockup "github.com/TheDevtop/exm/sti/drv-mockup"
)

func TestSearch(t *testing.T) {
	// Correct results
	good := []string{"This is the content of foobar"}

	// Setup mock driver
	var (
		scanPtr *bufio.Scanner
		err     error
	)
	if scanPtr, err = drvmockup.Stream("foobar"); err != nil {
		t.Fatal(err)
	}
	// Setup cache
	var rePtr *regexp.Regexp
	rec.Setup(false)
	if rePtr, err = rec.Receive("content"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := Search(rePtr, scanPtr); result[0] != good[0] {
		t.Fatal("Test result and correct answer mismatch!")
	}
}

func TestReplace(t *testing.T) {
	// Correct results
	good := []string{"This is John Cena of foobar"}

	// Setup mock driver
	var (
		scanPtr *bufio.Scanner
		err     error
	)
	if scanPtr, err = drvmockup.Stream("foobar"); err != nil {
		t.Fatal(err)
	}
	// Setup cache
	var rePtr *regexp.Regexp
	rec.Setup(false)
	if rePtr, err = rec.Receive("the content"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := Replace(rePtr, scanPtr, "John Cena"); result[0] != good[0] {
		t.Fatal("Test result and correct answer mismatch!")
	}
}

func TestMapReduce(t *testing.T) {
	// Correct results
	good := []string{"This is John Cena of foobar"}

	// Setup mock driver
	var (
		scanPtr *bufio.Scanner
		err     error
	)
	if scanPtr, err = drvmockup.Stream("foobar"); err != nil {
		t.Fatal(err)
	}
	// Setup cache
	var rePtr *regexp.Regexp
	rec.Setup(false)
	if rePtr, err = rec.Receive("the content"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := MapReduce(rePtr, scanPtr, "John Cena"); result[0] != good[0] {
		t.Fatal("Test result and correct answer mismatch!")
	}
}

func TestReduce(t *testing.T) {
	// Correct results
	good := []string{"This is John Cena of foobar"}

	// Setup mock driver
	var (
		scanPtr *bufio.Scanner
		err     error
	)
	if scanPtr, err = drvmockup.Stream("foobar"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := Reduce(scanPtr); result[0] != good[0] {
		t.Logf("%s\n", result[0])
		t.Fatal("Test result and correct answer mismatch!")
	}
}
