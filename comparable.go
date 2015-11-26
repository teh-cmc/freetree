package freetree

// -----------------------------------------------------------------------------

// Comparable exposes a single Less method.
type Comparable interface {
	// Less returns true if `c1` < `c2`
	Less(c2 Comparable) bool
}

// -----------------------------------------------------------------------------

// ComparableArray represents a sortable array of Comparables.
type ComparableArray []Comparable

// Len returns the length of the array.
func (ca ComparableArray) Len() int { return len(ca) }

// Less returns true if `ca[i]` < `ca[j]`.
func (ca ComparableArray) Less(i, j int) bool { return ca[i].Less(ca[j]) }

// Swap swaps the values of `ca[i]` and `ca[j]`.
func (ca ComparableArray) Swap(i, j int) { ca[i], ca[j] = ca[j], ca[i] }

// -----------------------------------------------------------------------------

// ComparableIterator is used to walk through a SimpleTree or a FreeTree.
type ComparableIterator func(c Comparable) bool
