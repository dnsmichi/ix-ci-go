package ix_ci

import (
	"reflect"
	"testing"
)

func TestSplitVersionStr(t *testing.T) {
	tables := []struct{
		v string
		r []string
	}{
		{"2.10.1", []string{ "2", "10", "1"} },
		{"1.14", []string{ "1", "14"} },
	}

	for _, table := range(tables) {
		res := SplitVersionStr(table.v)

		if reflect.DeepEqual(res, table.r) != true {
			t.Errorf("Split version of '%s' was incorrect, got array %s, want array %s.",
				table.v, res, table.r)
		}
	}
}