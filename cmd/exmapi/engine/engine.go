package engine

import (
	"bufio"
	"fmt"
	"io/fs"
	"regexp"

	"github.com/TheDevtop/exm/lib"
)

func SearchObject(restr, object string) ([]string, error) {
	var (
		stream  *bufio.Scanner
		err     error
		re      *regexp.Regexp
		results = make([]string, 0, 32)
	)

	if stream, err = getObject(object); err != nil {
		return nil, err
	}
	if re, err = getRegex(restr); err != nil {
		return nil, err
	}
	for stream.Scan() {
		if ln := stream.Text(); re.MatchString(ln) {
			results = append(results, ln)
		}
	}
	return results, nil
}

func SearchGlobal(restr string) ([]string, error) {
	var (
		stream  *bufio.Scanner
		err     error
		re      *regexp.Regexp
		tempMap = make(map[string]bool)
		results = make([]string, 0, 32)
	)

	if re, err = getRegex(restr); err != nil {
		return nil, err
	}
	for _, entry := range config.Sources {
		if stream, err = getObject(entry.Object); err != nil {
			return nil, err
		}
		for stream.Scan() {
			if ln := stream.Text(); re.MatchString(ln) {
				tempMap[entry.Object] = true
			}
		}
	}
	for key := range tempMap {
		results = append(results, key)
	}
	return results, nil
}

func IndexObject(object string) ([]string, error) {
	var (
		stream  *bufio.Scanner
		err     error
		results = make([]string, 0, 32)
		tempMap = make(map[string]bool)
	)

	if stream, err = getObject(object); err != nil {
		return nil, err
	}
	stream.Split(bufio.ScanWords)
	for stream.Scan() {
		tempMap[stream.Text()] = true
	}
	for str := range tempMap {
		results = append(results, str)
	}
	return results, nil
}

func IndexGlobal() []string {
	var results = make([]string, 0, 32)
	for _, entry := range config.Sources {
		results = append(results, entry.Object)
	}
	return results
}

func MetaObject(object string) lib.FormMetadata {
	var (
		fi  fs.FileInfo
		err error
		fm  = new(lib.FormMetadata)
		se  lib.SourceEntry
	)
	if fi, err = getMetadata(object); err != nil {
		fm.Error = err.Error()
		return *fm
	}
	for _, se = range config.Sources {
		if se.Object == object {
			goto RESULT // Considered harmful
		}
	}
	fm.Error = fmt.Sprintf(lib.ErrObject, "No object", object)
	return *fm

RESULT:
	fm.Error = ""
	fm.Object = object
	fm.Type = se.Type
	fm.Source = se.Address
	fm.LastModified = fi.ModTime().String()
	fm.Size = fi.Size()
	return *fm
}
