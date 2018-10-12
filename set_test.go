package setof

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewStringToString()
	set.Add("a", "a")
	set.Add("b", "b")
	set.Add("c", "c")
	if !reflect.DeepEqual(set.Keys(), []string{"a", "b", "c"}) {
		t.Errorf("Expected set to only contain a, b, c, got %v", set.Keys())
	}
	set.Add("a", "aa")
	aa, ok := set.Get("a")
	if !ok {
		t.Errorf("Expected to be able to get the value 'aa', but couldn't")
	}
	if aa != "aa" {
		t.Errorf("Expected to get value 'aa' but got '%v'", aa)
	}
	set.Del("c")
	if !reflect.DeepEqual(set.Keys(), []string{"a", "b"}) {
		t.Errorf("Expected set to only contain a, b, got %v", set.Keys())
	}
	set.Add("c", "123")
	if !reflect.DeepEqual(set.Keys(), []string{"a", "b", "c"}) {
		t.Errorf("Expected set to contain a, b, c after restore, got %v", set.Keys())
	}
	c, ok := set.Get("c")
	if !ok {
		t.Errorf("Expected to be able to get the value, but couldn't")
	}
	if c != "123" {
		t.Errorf("Expected to be able to get the value, got %v", c)
	}
	set.Del("b")
	if !reflect.DeepEqual(set.Keys(), []string{"a", "c"}) {
		t.Errorf("Expected set to only contain a, c, got %v", set.Keys())
	}
}
