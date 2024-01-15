//go:build linux

package flags

import (
	"slices"
	"sort"
	"testing"
)

func TestGetAllFlags(t *testing.T) {
	flags, err := GetAllFlags()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(flags)
}

func TestGetFlagByName(t *testing.T) {
	flags, err := GetFlagsByName("ens160")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(flags)
	expect := []string{
		"UP", "BROADCAST", "RUNNING", "MULTICAST",
	}
	sort.Strings(flags)
	sort.Strings(expect)
	if !slices.Equal(flags, expect) {
		t.Fatal("failed to get flags by name")
	}
}
