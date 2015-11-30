// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/teh-cmc/freetree"
)

// -----------------------------------------------------------------------------

/////
// Simple example of usage
//
//   go run examples/simple.go
//
/////

// Int is a simple integer that implements the Comparable interface.
type Int int

// Less returns true if `i1` < `i2`.
func (i1 Int) Less(i2 freetree.Comparable) bool { return i1 < i2.(Int) }

func main() {
	// build a new SimpleTree and insert 3 integers in it
	st := freetree.NewSimpleTree().Insert(Int(17), Int(66), Int(42))

	// print 42
	fmt.Println(st.Ascend(Int(42)))
	// print <nil>
	fmt.Println(st.Ascend(Int(43)))

	// print [42 17 66]
	fmt.Println(st.Flatten())

	// rebalance the tree
	st.Rebalance()

	// print [17 66 42]
	fmt.Println(st.Flatten())

	// build a new FreeTree using the data from the SimpleTree
	ft, err := freetree.NewFreeTree(st)
	if err != nil {
		log.Fatal(err)
	}

	// delete the SimpleTree and clear the garbage
	st = st.DeleteGC()

	// print 42
	fmt.Println(ft.Ascend(Int(42)))
	// print <nil>
	fmt.Println(ft.Ascend(Int(43)))

	// print [17 66 42]
	fmt.Println(ft.Flatten())

	// delete the FreeTree
	ft.Delete()
}
