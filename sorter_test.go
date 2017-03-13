package sorter_test

import (
	"sort"

	"fmt"

	"github.com/dave/sorter"
)

func ExampleNew() {

	type person struct{ name string }

	s := []person{{name: "foo"}, {name: "bar"}, {name: "baz"}, {name: "qux"}}
	sort.Sort(sorter.New(
		len(s),
		func(i, j int) { s[i], s[j] = s[j], s[i] },
		func(i, j int) bool { return s[i].name < s[j].name },
	))

	fmt.Println(s)
	// Output: [{bar} {baz} {foo} {qux}]

}
