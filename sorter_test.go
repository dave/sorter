package sorter_test

import (
	"sort"
	"testing"

	"reflect"

	"github.com/davelondon/sorter"
)

func TestNew(t *testing.T) {

	s := []person{{name: "foo"}, {name: "bar"}, {name: "baz"}}
	sort.Sort(sorter.New(
		len(s),
		func(i, j int) { s[i], s[j] = s[j], s[i] },
		func(i, j int) bool { return s[i].name < s[j].name },
	))

	expected := []person{{"bar"}, {"baz"}, {"foo"}}
	if !reflect.DeepEqual(s, expected) {
		t.Fatalf("Expected: %v. Got: %v", expected, s)
	}
}

type person struct {
	name string
}
