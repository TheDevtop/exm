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
	rec.Setup(2, false)
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
	rec.Setup(2, false)
	if rePtr, err = rec.Receive("the content"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := Replace(rePtr, scanPtr, "John Cena"); result[0] != good[0] {
		t.Fatal("Test result and correct answer mismatch!")
	}
}

func TestMapReduce(t *testing.T) {
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
	rec.Setup(2, false)
	if rePtr, err = rec.Receive("the content"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := MapReduce(rePtr, scanPtr, "John Cena"); result < 1 {
		t.Fatal("Test result should be higher then 0")
	}
}

func TestReduce(t *testing.T) {
	// Setup mock driver
	var (
		scanPtr *bufio.Scanner
		err     error
	)
	if scanPtr, err = drvmockup.Stream("foobar"); err != nil {
		t.Fatal(err)
	}
	// Test
	if result := Reduce(scanPtr); result[0] != "" {
		t.Logf("%s\n", result[0])
		t.Fatal("Test result and correct answer mismatch!")
	}
}

func TestMatch(t *testing.T) {
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
	rec.Setup(2, false)
	if rePtr, err = rec.Receive("foo*"); err != nil {
		t.Fatal(err)
	}

	// Test should work
	if !Match(rePtr, scanPtr) {
		t.Fatal("Test should match")
	}

	if rePtr, err = rec.Receive("NoMatch"); err != nil {
		t.Fatal(err)
	}

	// Test should not work
	if Match(rePtr, scanPtr) {
		t.Fatal("Test should not match")
	}
}
