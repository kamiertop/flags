//go:build linux

package flags

var m map[uint16]string

func FlagNumber() int {
	return 0
}

func FlagString() string {
	return ""
}

func FlagSlice() []string {
	return []string{}
}
