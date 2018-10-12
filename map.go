package setof

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "KeyType=string,int,int64 ValueType=string,int,int64"

import (
	"sort"
	"sync/atomic"

	"github.com/cheekybits/genny/generic"
)

// KeyType of the map.
type KeyType generic.Type

// ValueType of the map.
type ValueType generic.Type

// NewKeyTypeToValueType creates a new map.
func NewKeyTypeToValueType() *KeyTypeToValueType {
	return &KeyTypeToValueType{
		mapKeysToIndex: make(map[KeyType]*indexToKeyTypeWithValueValueType),
	}
}

// KeyTypeToValueType is a map which retains the order of the keys.
type KeyTypeToValueType struct {
	mapKeysToIndex map[KeyType]*indexToKeyTypeWithValueValueType
	index          int64
}

// Add an item to the map.
func (m *KeyTypeToValueType) Add(k KeyType, v ValueType) {
	if kv, ok := m.mapKeysToIndex[k]; ok {
		kv.value = v
		return
	}
	m.mapKeysToIndex[k] = &indexToKeyTypeWithValueValueType{
		index: atomic.AddInt64(&m.index, 1),
		key:   k,
		value: v,
	}
}

// Get an item from the map.
func (m *KeyTypeToValueType) Get(k KeyType) (v ValueType, ok bool) {
	kv, ok := m.mapKeysToIndex[k]
	v = kv.value
	return
}

// Del deletes an item from the map.
func (m *KeyTypeToValueType) Del(k KeyType) {
	delete(m.mapKeysToIndex, k)
}

// Keys returns all of the keys within the map.
func (m *KeyTypeToValueType) Keys() (keys []KeyType) {
	kvs := make(indexToKeyTypeWithValueValueTypes, len(m.mapKeysToIndex))
	var index int
	for _, kv := range m.mapKeysToIndex {
		kvs[index] = *kv
		index++
	}
	sort.Sort(kvs)
	keys = make([]KeyType, len(m.mapKeysToIndex))
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
