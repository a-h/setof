# setof

A package containing sets which retain the order of the inserted keys.

```go
func ExampleStringSet() {
	set := Strings("a", "b")
	set.Add("d")
	set.Add("c")
	set.Del("d")
	fmt.Println(set.Values())
	// Output: [a b c]
}
```
