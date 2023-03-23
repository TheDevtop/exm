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

// Search for pattern in object and replace with mapping
func Replace(rex *regexp.Regexp, streamer *bufio.Scanner, mapping string) []string {
	var result = make([]string, 0, 32)
	for streamer.Scan() {
		result = append(result, rex.ReplaceAllString(streamer.Text(), mapping))
	}
	return result
}

// Search for pattern in object and map the matches exclusively
func MapReduce(rex *regexp.Regexp, streamer *bufio.Scanner, mapping string) []string {
	var result = make([]string, 0, 32)
	for streamer.Scan() {
		if ln := streamer.Text(); rex.MatchString(ln) {
			result = append(result, rex.ReplaceAllString(ln, mapping))
		}
	}
	return result
}

// Reduce object to its dictionary
func Reduce(streamer *bufio.Scanner) []string {
	var dict = make(map[string]bool, 32)
	var result = make([]string, 32)

	streamer.Split(bufio.ScanWords)
	for streamer.Scan() {
		dict[streamer.Text()] = true
	}
	for key := range dict {
		result = append(result, key)
	}
	return result
}
