package setof

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int,int64 ValueType=string,int,int64"

import (
	"sort"
	"sync/atomic"

	"github.com/cheekybits/genny/generic"
)

// KeyType of the set.
type KeyType generic.Type

// ValueType of the set.
type ValueType generic.Type

// NewKeyTypeToValueType creates a new map of KeyType to ValueType.
func NewKeyTypeToValueType() *KeyTypeToValueType {
	return &KeyTypeToValueType{
		mapKeysToIndex: make(map[KeyType]*indexToKeyTypeWithValueValueType),
	}
}

// KeyTypeToValueType is a map of KeyType to ValueType which retains the order of the keys.
type KeyTypeToValueType struct {
	mapKeysToIndex map[KeyType]*indexToKeyTypeWithValueValueType
	index          int64
}

// Add an item to the set.
func (ss *KeyTypeToValueType) Add(k KeyType, v ValueType) {
	if kv, ok := ss.mapKeysToIndex[k]; ok {
		kv.value = v
		return
	}
	ss.mapKeysToIndex[k] = &indexToKeyTypeWithValueValueType{
		index: atomic.AddInt64(&ss.index, 1),
		key:   k,
		value: v,
	}
}

// Get an item from the set.
func (ss *KeyTypeToValueType) Get(k KeyType) (v ValueType, ok bool) {
	kv, ok := ss.mapKeysToIndex[k]
	v = kv.value
	return
}

// Del deletes an item from the set.
func (ss *KeyTypeToValueType) Del(k KeyType) {
	delete(ss.mapKeysToIndex, k)
}

// Keys returns all of the keys within the set.
func (ss *KeyTypeToValueType) Keys() (keys []KeyType) {
	kvs := make(indexToKeyTypeWithValueValueTypes, len(ss.mapKeysToIndex))
	var index int
	for _, kv := range ss.mapKeysToIndex {
		kvs[index] = *kv
		index++
	}
	sort.Sort(kvs)
	keys = make([]KeyType, len(ss.mapKeysToIndex))
	for i, v := range kvs {
		keys[i] = v.key
	}
	return keys
}

type indexToKeyTypeWithValueValueType struct {
	index int64
	key   KeyType
	value ValueType
}

type indexToKeyTypeWithValueValueTypes []indexToKeyTypeWithValueValueType

func (d indexToKeyTypeWithValueValueTypes) Len() int {
	return len(d)
}

func (d indexToKeyTypeWithValueValueTypes) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d indexToKeyTypeWithValueValueTypes) Less(i, j int) bool {
	return d[i].index < d[j].index
}
