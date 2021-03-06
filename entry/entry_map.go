package entry

import (
	"github.com/iancoleman/orderedmap"
)

// OrderedMap is an ordered map of entries.
type OrderedMap struct {
	orderedMap *orderedmap.OrderedMap
}

// NewOrderedMap creates a new OrderedMap of entries.
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		orderedMap: orderedmap.New(),
	}
}

// NewOrderedMapFromEntries creates a new OrderedMap of entries from a slice.
func NewOrderedMapFromEntries(entries []*Entry) *OrderedMap {
	orderedMap := NewOrderedMap()

	for _, e := range entries {
		if e == nil {
			continue
		}

		orderedMap.Set(e.Hash.String(), e)
	}

	return orderedMap
}

// Merge will fusion two OrderedMap of entries.
func (o *OrderedMap) Merge(other *OrderedMap) *OrderedMap {
	newMap := o.Copy()

	otherKeys := other.Keys()
	for _, k := range otherKeys {
		val, _ := other.Get(k)
		newMap.Set(k, val)
	}

	return newMap
}

// Copy creates a copy of an OrderedMap.
func (o *OrderedMap) Copy() *OrderedMap {
	newMap := NewOrderedMap()
	keys := o.Keys()

	for _, k := range keys {
		val, _ := o.Get(k)
		newMap.Set(k, val)
	}

	return newMap
}

// Get retrieves an Entry using its key.
func (o *OrderedMap) Get(key string) (*Entry, bool) {
	val, exists := o.orderedMap.Get(key)
	entry, ok := val.(*Entry)
	if !ok {
		exists = false
	}

	return entry, exists
}

// UnsafeGet retrieves an Entry using its key, returns nil if not found.
func (o *OrderedMap) UnsafeGet(key string) *Entry {
	val, _ := o.Get(key)

	return val
}

// Set defines an Entry in the map for a given key.
func (o *OrderedMap) Set(key string, value *Entry) {
	o.orderedMap.Set(key, value)
}

// Slice returns an ordered slice of the values existing in the map.
func (o *OrderedMap) Slice() []*Entry {
	out := []*Entry{}

	keys := o.orderedMap.Keys()
	for _, k := range keys {
		out = append(out, o.UnsafeGet(k))
	}

	return out
}

// Delete removes an Entry from the map for a given key.
func (o *OrderedMap) Delete(key string) {
	o.orderedMap.Delete(key)
}

// Keys retrieves the ordered list of keys in the map.
func (o *OrderedMap) Keys() []string {
	return o.orderedMap.Keys()
}

// SortKeys orders the map keys using your sort func.
func (o *OrderedMap) SortKeys(sortFunc func(keys []string)) {
	o.orderedMap.SortKeys(sortFunc)
}

// Sort orders the map using your sort func.
func (o *OrderedMap) Sort(lessFunc func(a *orderedmap.Pair, b *orderedmap.Pair) bool) {
	o.orderedMap.Sort(lessFunc)
}

// Len gets the length of the map.
func (o *OrderedMap) Len() int {
	return len(o.orderedMap.Keys())
}

// At gets an item at the given index in the map, returns nil if not found.
func (o *OrderedMap) At(index uint) *Entry {
	keys := o.Keys()

	if uint(len(keys)) < index {
		return nil
	}

	return o.UnsafeGet(keys[index])
}
