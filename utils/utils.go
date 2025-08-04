// Package utils contains utility functions for various operations.
package utils

import (
	"slices"
	"sort"
)

// EqualUnordered checks if two slices contain the same elements, regardless of order.
func EqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := slices.Clone(a)
	bCopy := slices.Clone(b)
	sort.Ints(aCopy)
	sort.Ints(bCopy)
	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}
	return true
}
