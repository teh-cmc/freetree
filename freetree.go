package freetree

// -----------------------------------------------------------------------------

// SimpleTree implements a simple binary search tree.
type SimpleTree struct {
	root *simpleNode
}

// NewSimpleTree returns an empty SimpleTree.
func NewSimpleTree() *SimpleTree {
	return &SimpleTree{}
}

// Insert inserts the given Comparables in the tree.
// It does not rebalance the tree, use Rebalance() for that.
//
// If the tree is currently empty and the passed-in Comparable are already
// sorted in increasing order, the tree will be perfectly balanced.
// This means you don't have to Rebalance() the tree if you've inserted all
// your pre-sorted data in one Insert() call.
func (st *SimpleTree) Insert(cs ...Comparable) {
	st.insert(cs)
}

// InsertArray is a helper to use Insert() with a ComparableArray.
func (st *SimpleTree) InsertArray(cs ComparableArray) {
	st.insert(cs)
}

func (st *SimpleTree) insert(cs ComparableArray) {
	l := len(cs)
	if l == 0 {
		return
	}

	st.root = st.root.insert(cs[l/2])

	if l > 1 {
		st.insert(cs[:l/2])
		st.insert(cs[l/2:])
	}
}

// Ascend returns the first element in the tree that is >= `pivot`.
func (st SimpleTree) Ascend(pivot Comparable) Comparable {
	return st.ascend(pivot)
}

func (st SimpleTree) ascend(pivot Comparable) Comparable {
	return st.root.ascend(pivot)
}

// -----------------------------------------------------------------------------

type simpleNode struct {
	left, right *simpleNode
	data        Comparable
}

func (sn *simpleNode) insert(c Comparable) *simpleNode {
	if sn == nil {
		return &simpleNode{data: c}
	}

	if c.Less(sn.data) {
		sn.left = sn.left.insert(c)
	} else {
		sn.right = sn.right.insert(c)
	}

	return sn
}

func (sn *simpleNode) ascend(pivot Comparable) Comparable {
	if sn == nil {
		return nil
	}

	if pivot.Less(sn.data) {
		return sn.data
	} else {
		return sn.right.ascend(pivot)
	}
}
