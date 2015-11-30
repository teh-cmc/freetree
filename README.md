# FreeTree ![Status](https://img.shields.io/badge/status-stable-green.svg?style=plastic) [![Build Status](http://img.shields.io/travis/teh-cmc/freetree.svg?style=plastic)](https://travis-ci.org/teh-cmc/freetree) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=plastic)](http://godoc.org/github.com/teh-cmc/freetree)

Generic binary search tree with zero GC overhead, for golang. Built on [mmm](https://github.com/teh-cmc/mmm).

## What you should know

`FreeTree` was developed mainly as a proof-of-concept for [mmm](https://github.com/teh-cmc/mmm); that is, to demonstrate how you can use `mmm` to avoid GC overhead in "pointer-heavy" Go software, without modifying nor complexifying your original design (i.e. entirely redesigning your software to avoid the use of pointers, which often leads to overly complex and less maintainable code).

Although I do use it for some big immutable caches of mine, `FreeTree`'s API is quite incomplete and could certainly be better designed; especially during initialization where a lot of unnecessary copying could probably be avoided.
Feel free to improve it :)

## Install

```bash
go get -u github.com/teh-cmc/freetree
```

## Example

Here's a simple example of usage (code [here](examples/simple.go)):

```Go
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
```

## Demonstration

Complete code for the following demonstration is available [here](experiment/experiment.go).

All of the results shown below were computed using a DELL XPS 15-9530 (i7-4712HQ@2.30GHz).

#### Case A: normal BST, 10 million integers

Let's look at GC performances while storing 10 million integers in a "classic" binary search tree:

```Go
// build 10 million integers
ints := make([]freetree.Comparable, 10*1e6)
// init our integers
for i := range ints {
	ints[i] = Int(i)
}

// build a new BST and insert our 10 million integers in it
// our integers are pre-sorted, so the tree will be perfectly balanced (because
// that's how FreeTree's insert API works)
st := freetree.NewSimpleTree().InsertArray(ints)

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected, and to prevent them from being optimized away
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, st.Ascend(Int(i*1e4)))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (normal BST, 10 million integers): %d us\n", (time.Now().UnixNano()-now)/1e3)
}
```

This prints:

```
value @ index 0: 0
GC time (normal BST, 10 million integers): 276821 us
value @ index 10000: 10000
GC time (normal BST, 10 million integers): 278205 us
value @ index 20000: 20000
GC time (normal BST, 10 million integers): 286721 us
value @ index 30000: 30000
GC time (normal BST, 10 million integers): 277796 us
value @ index 40000: 40000
GC time (normal BST, 10 million integers): 277528 us
```

That's an average ~278ms per GC call.
Let's move to case B where we'll store 10 million integers in a FreeTree.

#### Case B: FreeTree, 10 million integers

Let's look at GC performances while storing 10 million integers in a `FreeTree` (i.e. a binary search tree with no GC overhead):

```Go
// build a new FreeTree using the data from our SimpleTree
ft, err := freetree.NewFreeTree(st)
if err != nil {
	log.Fatal(err)
}

// delete the SimpleTree and get rid of the garbage
st = st.DeleteGC()

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, ft.Ascend(Int(i*1e4)))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (FreeTree, 10 million integers): %d us\n", (time.Now().UnixNano()-now)/1e3)
}
```

This prints:

```
value @ index 0: 0
GC time (FreeTree, 10 million integers): 3287 us
value @ index 10000: 10000
GC time (FreeTree, 10 million integers): 3527 us
value @ index 20000: 20000
GC time (FreeTree, 10 million integers): 3346 us
value @ index 30000: 30000
GC time (FreeTree, 10 million integers): 3447 us
value @ index 40000: 40000
GC time (FreeTree, 10 million integers): 3104 us
```

We went from a ~278ms average to a ~3.3ms average; and most importantly, we did that without modifying our design: internally, both trees' structures and `Ascend` API work the same way.


## License ![License](https://img.shields.io/badge/license-MIT-blue.svg?style=plastic)

The MIT License (MIT) - see LICENSE for more details

Copyright (c) 2015	Clement 'cmc' Rey	<cr.rey.clement@gmail.com>
