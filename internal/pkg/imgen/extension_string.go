// Code generated by "stringer -type Extension extension.go"; DO NOT EDIT.

package imgen

import "strconv"

const _Extension_name = "PNGJPEGGIF"

var _Extension_index = [...]uint8{0, 3, 7, 10}

func (i Extension) String() string {
	if i >= Extension(len(_Extension_index)-1) {
		return "Extension(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Extension_name[_Extension_index[i]:_Extension_index[i+1]]
}