package freetree

import (
	"runtime"
	"runtime/debug"
	"sort"
)

// -----------------------------------------------------------------------------

// SimpleTree implements a simple binary search tree.
type SimpleTree struct {
	root  *simpleNode
	nodes uint
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
func (st *SimpleTree) Insert(cs ...Comparable) *SimpleTree {
	return st.InsertArray(cs)
}

// InsertArray is a helper to use Insert() with a ComparableArray.
func (st *SimpleTree) InsertArray(ca ComparableArray) *SimpleTree {
	st.insert(ca)
	st.nodes += uint(len(ca))

	return st
}

func (st *SimpleTree) insert(ca ComparableArray) {
	l := len(ca)
	if l == 0 {
		return
	}

	st.root = st.root.insert(ca[l/2])

	if l > 1 {
		st.insert(ca[:l/2])
		st.insert(ca[l/2+1:])
	}
}

// Ascend returns the first element in the tree that is == `pivot`.
func (st SimpleTree) Ascend(pivot Comparable) Comparable {
	return st.ascend(pivot)
}

func (st SimpleTree) ascend(pivot Comparable) Comparable {
	return st.root.ascend(pivot)
}

// Rebalance rebalances the tree to guarantee O(log(n)) search complexity.
//
// Rebalancing is implemented as straightforwardly as possible: it's dumb.
// I strongly suggest running the garbage collector and scavenger once it's done.
//   runtime.GC()
//   debug.FreeOSMemory()
func (st *SimpleTree) Rebalance() *SimpleTree {
	flat := st.flatten()
	sort.Sort(flat)

	st.root = nil
	st.insert(flat)

	return st
}

// Rebalance rebalances the tree and runs the garbage collector.
func (st *SimpleTree) RebalanceGC() *SimpleTree {
	st.Rebalance()
	runtime.GC()
	debug.FreeOSMemory()

	return st
}

// Flatten returns the content of the tree as a ComparableArray.
func (st SimpleTree) Flatten() ComparableArray {
	return st.flatten()
}

func (st SimpleTree) flatten() ComparableArray {
	ca := make(ComparableArray, 0, st.nodes)
	return st.root.flatten(ca)
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
		return sn.left.ascend(pivot)
	} else if sn.data.Less(pivot) {
		return sn.right.ascend(pivot)
	}

	return sn.data
}

func (sn *simpleNode) flatten(ca ComparableArray) ComparableArray {
	if sn == nil {
		return ca
	}

	ca = sn.left.flatten(ca)
	ca = sn.right.flatten(ca)

	return append(ca, sn.data)
}
