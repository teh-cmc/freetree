package freetree

import (
	"unsafe"

	"github.com/teh-cmc/mmm"
)

// -----------------------------------------------------------------------------

// FreeTree implements a binary search tree with zero GC overhead.
type FreeTree struct {
	nodeChunk mmm.MemChunk
	dataChunk mmm.MemChunk
	root      *freeNode
}

// NewFreeTree returns a new FreeTree using the data of a supplied SimpleTree.
func NewFreeTree(st *SimpleTree) (*FreeTree, error) {
	nbNodes := st.nodes
	nodeChunk, err := mmm.NewMemChunk(freeNode{}, nbNodes)
	if err != nil {
		return nil, err
	}
	dataChunk, err := mmm.NewMemChunk(st.root.data, nbNodes)
	if err != nil {
		return nil, err
	}

	ft := &FreeTree{nodeChunk: nodeChunk, dataChunk: dataChunk}
	for _, n := range st.flattenNodes() {
		node := (*freeNode)(ft.nodeChunk.Index(int(n.id)))
		node.id = n.id
		if n.left != nil {
			node.left = uintptr(nodeChunk.Index(int(n.left.id)))
		}
		if n.right != nil {
			node.right = uintptr(nodeChunk.Index(int(n.right.id)))
		}
		node.data = uintptr(dataChunk.Index(int(n.id)))

		if n == st.root {
			ft.root = node
		}
	}

	return ft, nil
}

// Ascend returns the first element in the tree that is == `pivot`.
func (ft FreeTree) Ascend(pivot Comparable) Comparable {
	return ft.ascend(pivot)
}

func (ft FreeTree) ascend(pivot Comparable) Comparable {
	return ft.root.ascend(pivot)
}

// Flatten returns the content of the tree as a ComparableArray.
func (ft FreeTree) Flatten() ComparableArray {
	return ft.flatten()
}

func (ft FreeTree) flatten() ComparableArray {
	ca := make(ComparableArray, 0, ft.nodeChunk.NbObjects())
	return ft.root.flatten(ca)
}

// -----------------------------------------------------------------------------

type freeNode struct {
	id          uint
	left, right uintptr
	data        uintptr
}

func (sn *freeNode) ascend(pivot Comparable) Comparable {
	if sn == nil {
		return nil
	}

	if pivot.Less(*(*Comparable)(unsafe.Pointer(sn.data))) {
		return ((*freeNode)(unsafe.Pointer(sn.left))).ascend(pivot)
	} else if (*(*Comparable)(unsafe.Pointer(sn.data))).Less(pivot) {
		return ((*freeNode)(unsafe.Pointer(sn.right))).ascend(pivot)
	}

	return *((*Comparable)(unsafe.Pointer(sn.data)))
}

func (sn *freeNode) flatten(ca ComparableArray) ComparableArray {
	if sn == nil {
		return ca
	}

	ca = ((*freeNode)(unsafe.Pointer(sn.left))).flatten(ca)
	ca = ((*freeNode)(unsafe.Pointer(sn.right))).flatten(ca)

	return append(ca, *((*Comparable)(unsafe.Pointer(sn.data))))
}
