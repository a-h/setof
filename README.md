# setof

A package containing maps and sets which retain the order of the inserted keys.

```go
func ExampleStringSet() {
	set := Strings("a", "b")
	set.Add("d")
	set.Add("c")
	set.Del("d")
	fmt.Println(set.Values())
	// Output: [a b c]
}

func ExampleStringToString() {
	m := NewStringToString()
	m.Add("a", "aa")
	m.Add("b", "bb")
	m.Add("b", "bbb")
	m.Add("c", "c")
	m.Del("c")
	value, ok := m.Get("b")
	fmt.Println(value, ok, m.Keys())
	// Output: bbb true [a b]
}
```
