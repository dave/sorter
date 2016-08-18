# sorter
Simplify sorting in Go

# Usage
```
func main() {
  s := []person{{name: "foo"}, {name: "bar"}, {name: "baz"}}
  sort.Sort(sorter.New(
  	len(s),
  	func(i, j int) { s[i], s[j] = s[j], s[i] },
  	func(i, j int) bool { return s[i].name < s[j].name },
  ))
}
type person struct {
	name string
}
```
