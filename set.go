package setof

//go:generate genny -in=$GOFILE -out=gen-$GOFILE gen "SetType=string,int,int64"

import (
	"encoding/json"
	"sort"
	"sync/atomic"

	"github.com/cheekybits/genny/generic"
)

// SetType of the set.
type SetType generic.Type

// SetTypes creates a new set of SetTypes.
func SetTypes(values ...SetType) *SetTypeSet {
	ss := &SetTypeSet{
		mapKeysToIndex: make(map[SetType]int64),
	}
	for _, v := range values {
		ss.Add(v)
	}
	return ss
}

// SetTypeSet is a set of SetTypes which retains the order that the keys were added.
type SetTypeSet struct {
	mapKeysToIndex map[SetType]int64
	index          int64
}

// Add an item to the set.
func (s *SetTypeSet) Add(v SetType) {
	if _, ok := s.mapKeysToIndex[v]; ok {
		return
	}
	s.mapKeysToIndex[v] = atomic.AddInt64(&s.index, 1)
}

// Contains determines whether an item is in the set.
func (s *SetTypeSet) Contains(v SetType) (ok bool) {
	_, ok = s.mapKeysToIndex[v]
	return
}

// Del deletes an item from the set.
func (s *SetTypeSet) Del(v SetType) {
	delete(s.mapKeysToIndex, v)
}

// Values returns all of the values within the set.
func (s *SetTypeSet) Values() (v []SetType) {
	values := make(indexToSetTypeValues, len(s.mapKeysToIndex))
	var index int
	for k, v := range s.mapKeysToIndex {
		values[index] = indexToSetTypeValue{
			index: v,
			value: k,
		}
		index++
	}
	sort.Sort(values)
	v = make([]SetType, len(s.mapKeysToIndex))
	for i, vv := range values {
		v[i] = vv.value
	}
	return
}

type indexToSetTypeValue struct {
	index int64
	value SetType
}

type indexToSetTypeValues []indexToSetTypeValue

func (d indexToSetTypeValues) Len() int {
	return len(d)
}

func (d indexToSetTypeValues) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d indexToSetTypeValues) Less(i, j int) bool {
	return d[i].index < d[j].index
}

// MarshalJSON marshals to JSON.
func (s *SetTypeSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON marshals from JSON.
func (s *SetTypeSet) UnmarshalJSON(data []byte) error {
	var a []SetType
	err := json.Unmarshal(data, &a)
	if err != nil {
		return err
	}
	(*s) = (*SetTypes(a...))
	return nil
}
