package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"regexp"
)

// simple regex based extraction of interfaces. heuristic, not
// guaranteed to work all the time. Might be confounded by
// comments, etc. But gives us a boost 99% of the time.
var regexIface = regexp.MustCompile(`[\t ]*type[\t ]+(\S+)[\t ]+interface[\t ]*\{`)

// return value m maps from interface name -> file it came from
func ScanFilesForInterfaces(fns []string) (m map[string]string, err error) {
	m = make(map[string]string)
	for _, fn := range fns {
		by, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, fmt.Errorf("error reading file '%s': '%v'", fn, err)
		}

		m0, err := FindInterfaces(bytes.NewBuffer(by), fn)
		if err != nil {
			// merge into m
			for k, v := range m0 {
				m[k] = v
			}
		}
	}
	return m, nil
}

func FindInterfaces(r io.Reader, fn string) (m map[string]string, err error) {
	m = make(map[string]string)
	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(r)
	lineno := 0
	for scanner.Scan() {
		lineno++
		line := scanner.Text()
		//fmt.Printf("checking line: '%s'\n", line)
		line = strings.TrimSpace(line)
		if len(line) <= 0 {
			continue
		}
		slc := regexIface.FindStringSubmatch(line)
		//fmt.Printf("slc ='%#v'\n", slc)
		if len(slc) >= 2 {
			m[slc[1]] = fn
		}
	}

	// check for errors
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return m, nil
}
