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
