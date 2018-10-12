package setof

import (
	"reflect"
	"sort"
	"sync/atomic"
)

// Strings provides a set of strings to interfaces which retains the order of the keys.
func Strings(s ...string) *StringSet {
	ss := NewStringSet()
	for _, v := range s {
		ss.Add(v, nil)
	}
	return ss
}

// NewStringSet creates a new map of strings to interfaces.
func NewStringSet() *StringSet {
	return &StringSet{
		mapKeysToIndex: make(map[string]keyValue),
	}
}

// StringSet is a map of strings to interface values which retains the order of the keys.
type StringSet struct {
	mapKeysToIndex map[string]keyValue
	index          int64
}

// Add an item to the set.
func (ss *StringSet) Add(s string, v interface{}) {
	if kv, ok := ss.mapKeysToIndex[s]; ok {
		kv.value = reflect.ValueOf(v)
		return
	}
	ss.mapKeysToIndex[s] = keyValue{
		index: atomic.AddInt64(&ss.index, 1),
		key:   s,
		value: v,
	}
}

// Get an item from the set.
func (ss *StringSet) Get(s string) (v interface{}, ok bool) {
	kv, ok := ss.mapKeysToIndex[s]
	v = kv.value
	return
}

// Del deletes an item from the set.
func (ss *StringSet) Del(s string) {
	delete(ss.mapKeysToIndex, s)
}

// Keys returns all of the keys within the set.
func (ss *StringSet) Keys() (keys []string) {
	kvs := make(keyValues, len(ss.mapKeysToIndex))
	var index int
	for _, kv := range ss.mapKeysToIndex {
		kvs[index] = kv
		index++
	}
	sort.Sort(kvs)
	keys = make([]string, len(ss.mapKeysToIndex))
	for i, v := range kvs {
		keys[i] = v.key
	}
	return keys
}

type keyValue struct {
	index int64
	key   string
	value interface{}
}

type keyValues []keyValue

func (d keyValues) Len() int {
	return len(d)
}

func (d keyValues) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d keyValues) Less(i, j int) bool {
	return d[i].index < d[j].index
}
