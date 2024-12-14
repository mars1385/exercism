package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	items map[string]bool
}

func New() Set {
	var set Set = Set{items: map[string]bool{}}

	return set
}

func NewFromSlice(l []string) Set {
	var set Set = Set{items: map[string]bool{}}

	for _, v := range l {
		if !set.Has(v) {

			set.items[v] = true
		}
	}
	return set
}

func (s Set) String() string {

	result := ""

	for v := range s.items {
		if len(result) > 0 {

			result += ", "
		}

		result += `"` + v + `"`

	}
	return "{" + result + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.items) <= 0
}

func (s Set) Has(elem string) bool {
	return s.items[elem]
}

func (s Set) Add(elem string) {

	if !s.Has(elem) {
		s.items[elem] = true
	}

}

func Subset(s1, s2 Set) bool {

	if s1.IsEmpty() {
		return true
	}
	if s2.IsEmpty() {
		return false
	}
	for k := range s1.items {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {

	if s1.IsEmpty() || s2.IsEmpty() {
		return true
	}
	for k := range s1.items {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}

	for k := range s1.items {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	var set Set = Set{items: map[string]bool{}}
	if s1.IsEmpty() || s2.IsEmpty() {
		return set
	}

	for k := range s1.items {
		if s2.Has(k) {
			set.items[k] = true
		}
	}
	return set
}

func Difference(s1, s2 Set) Set {
	var set Set = Set{items: map[string]bool{}}
	if s1.IsEmpty() && s2.IsEmpty() {
		return set
	}

	for k := range s1.items {
		if !s2.Has(k) {
			set.items[k] = true
		}
	}
	return set
}

func Union(s1, s2 Set) Set {
	var set Set = Set{items: map[string]bool{}}
	if s1.IsEmpty() && s2.IsEmpty() {
		return set
	}

	for k := range s1.items {

		set.items[k] = true

	}

	for k := range s2.items {

		set.items[k] = true

	}
	return set
}
