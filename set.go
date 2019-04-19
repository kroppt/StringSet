package set

import (
	"sort"
	"strings"
)

// Set is an string set
type Set struct {
	m map[string]struct{}
}

// NewSet returns a pointer to an empty set
func NewSet() Set {
	m := make(map[string]struct{})
	return Set{m}
}

// NewSetInit returns a pointer to a set with the passed string
func NewSetInit(k string) (ns Set) {
	ns = NewSet()
	ns.add(k)
	return ns
}

// Add inserts the given string into the set if it does not already exist
func (s Set) Add(k string) {
	if s.Contains(k) {
		return
	}
	s.add(k)
}

// Remove deletes the given string from the set if it exists
func (s Set) Remove(k string) {
	if !s.Contains(k) {
		return
	}
	delete(s.m, k)
}

// Union merges the second set onto the first and returns whether the set changed
func (s Set) Union(os Set) bool {
	b := false
	for i := range os.m {
		if s.Contains(i) {
			continue
		}
		s.add(i)
		b = true
	}
	return b
}

// Copy returns a duplicate set
func (s Set) Copy() (ns Set) {
	ns = NewSet()
	for i := range s.m {
		ns.add(i)
	}
	return ns
}

// Range executes the ranging function for each element
func (s Set) Range(f func(string) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}

// Contains returns whether the set contains the string
func (s Set) Contains(k string) bool {
	_, ok := s.m[k]
	return ok
}

// IsEmpty returns true if the set size is zero
func (s Set) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns the number of items in the set
func (s Set) Size() int {
	return len(s.m)
}

// Equals returns whether the sets are equal
func (s Set) Equals(os Set) bool {
	if s.Size() != os.Size() {
		return false
	}
	equal := true
	s.Range(func(k string) bool {
		if !os.Contains(k) {
			equal = false
			return false
		}
		return true
	})
	return equal
}

// Print returns a string representation of the set
func (s Set) Print() string {
	strs := make([]string, 0)
	for k := range s.m {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	var sb strings.Builder
	sb.WriteString("{ ")
	for _, str := range strs {
		sb.WriteString(str)
		sb.WriteRune(' ')
	}
	sb.WriteString("}")
	return sb.String()
}

func (s Set) add(k string) {
	s.m[k] = struct{}{}
}
