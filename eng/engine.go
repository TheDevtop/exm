package eng

import (
	"bufio"
	"regexp"

	"github.com/TheDevtop/exm/conio"
)

// Search for pattern in object
func Search(pattern string, streamer *bufio.Scanner) ([]string, error) {
	const fprobe = "eng.Search"
	var result = make([]string, 0, 18)
	if re, err := regexp.Compile(pattern); err != nil {
		conio.Probeln(fprobe, err.Error())
		return nil, err
	} else {
		for streamer.Scan() {
			if ln := streamer.Text(); re.MatchString(ln) {
				result = append(result, ln)
			}
		}
		return result, nil
	}
}
