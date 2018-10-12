package setof

import (
	"fmt"
	"reflect"
	"testing"
)

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

func TestSet(t *testing.T) {
	set := Strings("a", "b", "c")
	if !reflect.DeepEqual(set.Values(), []string{"a", "b", "c"}) {
		t.Errorf("Expected set to only contain a, b, c, got %v", set.Values())
	}
	set.Add("a")
	ok := set.Contains("a")
	if !ok {
		t.Errorf("Expected to be able to get the value 'a', but couldn't")
	}
	set.Del("c")
	if !reflect.DeepEqual(set.Values(), []string{"a", "b"}) {
		t.Errorf("Expected set to only contain a, b, got %v", set.Values())
	}
	set.Add("c")
	if !reflect.DeepEqual(set.Values(), []string{"a", "b", "c"}) {
		t.Errorf("Expected set to contain a, b, c after restore, got %v", set.Values())
	}
	set.Del("b")
	if !reflect.DeepEqual(set.Values(), []string{"a", "c"}) {
		t.Errorf("Expected set to only contain a, c, got %v", set.Values())
	}
}

func TestStringToStringMap(t *testing.T) {
	m := NewStringToString()
	m.Add("a", "aa")
	m.Add("b", "bb")
	m.Add("c", "cc")
	if !reflect.DeepEqual(m.Keys(), []string{"a", "b", "c"}) {
		t.Errorf("Expected map to only contain a, b, c, got %v", m.Keys())
	}
	m.Add("a", "aaa")
	aaa, ok := m.Get("a")
	if !ok {
		t.Errorf("Expected to be able to get 'a', but couldn't")
	}
	if aaa != "aaa" {
		t.Errorf("Expected to get value 'aaa' but got '%v'", aaa)
	}
	m.Del("c")
	if !reflect.DeepEqual(m.Keys(), []string{"a", "b"}) {
		t.Errorf("Expected m to only contain a, b, got %v", m.Keys())
	}
	m.Add("c", "123")
	if !reflect.DeepEqual(m.Keys(), []string{"a", "b", "c"}) {
		t.Errorf("Expected m to contain a, b, c after restore, got %v", m.Keys())
	}
	c, ok := m.Get("c")
	if !ok {
		t.Errorf("Expected to be able to get the value, but couldn't")
	}
	if c != "123" {
		t.Errorf("Expected to be able to get the value, got %v", c)
	}
	m.Del("b")
	if !reflect.DeepEqual(m.Keys(), []string{"a", "c"}) {
		t.Errorf("Expected map to only contain a, c, got %v", m.Keys())
	}
}
