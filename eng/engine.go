package eng

import (
	"bufio"
	"regexp"
)

// Search for pattern in object
func Search(rex *regexp.Regexp, streamer *bufio.Scanner) []string {
	var result = make([]string, 0, 32)
	for streamer.Scan() {
		if ln := streamer.Text(); rex.MatchString(ln) {
			result = append(result, ln)
		}
	}
	return result
}
