package main

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/teh-cmc/freetree"
)

// -----------------------------------------------------------------------------

/////
// This file shows how to use a FreeTree and how it affects GC performances
// compared to a classic binary search tree.
//
//   go run experiment.go
//
/////

// Int is a simple integer that implements the Comparable interface.
type Int int

// Less returns true if `i1` < `i2`.
func (i1 Int) Less(i2 freetree.Comparable) bool { return i1 < i2.(Int) }

func main() {

	// build 10 million integers
	ints := make([]freetree.Comparable, 10*1e6)
	// init our integers
	for i := range ints {
		ints[i] = Int(i)
	}

	////////////////////////////////////////
	// A: Normal BST, 10 million integers //
	////////////////////////////////////////
	fmt.Println(`Case A: GC performances while storing 10 million integers in a classic binary search tree` + "\n")

	// build a new BST and insert our 10 million integers in it
	// our integers are pre-sorted, so the tree will be perfectly balanced
	st := freetree.NewSimpleTree().InsertArray(ints)

	// get rid of init garbage
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected, and to prevent them from being optimized away
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, st.Ascend(Int(i*1e4)))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (normal BST, 10 million integers): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (normal BST, 10 million integers): 276821 us
	//   value @ index 10000: 10000
	//   GC time (normal BST, 10 million integers): 278205 us
	//   value @ index 20000: 20000
	//   GC time (normal BST, 10 million integers): 286721 us
	//   value @ index 30000: 30000
	//   GC time (normal BST, 10 million integers): 277796 us
	//   value @ index 40000: 40000
	//   GC time (normal BST, 10 million integers): 277528 us

	fmt.Println()

	//////////////////////////////////////
	// B: FreeTree, 10 million integers //
	//////////////////////////////////////
	fmt.Println(`Case B: GC performances while storing 10 million integers in a FreeTree (i.e. a binary search tree with no GC overhead)` + "\n")

	// build a new FreeTree using the data from our SimpleTree
	ft, err := freetree.NewFreeTree(st)
	if err != nil {
		log.Fatal(err)
	}

	// delete the SimpleTree
	st = st.Delete()
	// get rid of (almost all) previous garbage
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, ft.Ascend(Int(i*1e4)))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (FreeTree, 10 million integers): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (FreeTree, 10 million integers): 3287 us
	//   value @ index 10000: 10000
	//   GC time (FreeTree, 10 million integers): 3527 us
	//   value @ index 20000: 20000
	//   GC time (FreeTree, 10 million integers): 3346 us
	//   value @ index 30000: 30000
	//   GC time (FreeTree, 10 million integers): 3447 us
	//   value @ index 40000: 40000
	//   GC time (FreeTree, 10 million integers): 3104 us
}
