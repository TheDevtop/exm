package rec

import "regexp"

var reCache map[string]*regexp.Regexp

func Setup() {
	reCache = make(map[string]*regexp.Regexp, 16)
}

func Generate(restr string) (*regexp.Regexp, error) {
	if re, ok := reCache[restr]; ok {
		return re, nil
	}
	if re, err := regexp.Compile(restr); err != nil {
		return nil, err
	} else {
		reCache[restr] = re
		return re, nil
	}
}
