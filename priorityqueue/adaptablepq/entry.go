package adaptablepq

import "fmt"

type Entry[K, V any] struct {
	Index int
	Key   K
	Value V
}

// String returns the string representation of the entry.
func (e *Entry[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", e.Key, e.Value)
}
