package msgp

import (
	"fmt"
	"strconv"
	"strings"
)

func Clue2Field(name string, clue string, zid int64) string {
	if zid >= 0 {
		return fmt.Sprintf("%s_zid%02d_%s", name, zid, clue)
	}
	// handle the missing zid, and don't write -1 as the zid.
	return fmt.Sprintf("%s__%s", name, clue)
}

func Field2Clue(fieldname string) (name string, clue string, zid int64, err error) {
	parts := strings.Split(fieldname, "_")
	n := len(parts)
	if n < 3 {
		err = fmt.Errorf("too few underscore (expect at least two) in fieldname '%s'", fieldname)
		return
	}
	clue = parts[n-1]
	if strings.HasPrefix(parts[n-2], "zid") {
		tmp, err2 := strconv.Atoi(parts[n-2][3:])
		if err2 == nil {
			zid = int64(tmp)
		} else {
			err = fmt.Errorf("problem parsing out _zid field in fieldname '%s': '%v'", fieldname, err2)
			return
		}
	}
	used := len(parts[n-1]) + len(parts[n-2]) + 2
	name = fieldname[:len(fieldname)-used]
	return
}
