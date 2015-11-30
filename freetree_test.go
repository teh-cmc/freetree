// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package freetree

import (
	"fmt"
	"log"
	"testing"
)

// -----------------------------------------------------------------------------

func TestFreeTree_pair_sorted_input(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(1), intTest(2), intTest(3), intTest(4), intTest(5), intTest(6)}
	st.InsertArray(cs)

	ft, err := NewFreeTree(st)
	if err != nil {
		t.Error(err)
	}
	flat := ft.Flatten()

	for i := range expected {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if ft.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(7)) != nil {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(8)) != nil {
		t.Error("unexpected retval")
	}
}

func TestFreeTree_pair_sorted_input_rebalanced(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(1), intTest(2), intTest(3), intTest(4), intTest(5), intTest(6)}

	st.InsertArray(cs)

	ft, err := NewFreeTree(st.Rebalance())
	if err != nil {
		t.Error(err)
	}
	flat := ft.Flatten()

	for i := range flat {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if ft.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(7)) != nil {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(8)) != nil {
		t.Error("unexpected retval")
	}
}

func TestFreeTree_pair_unsorted_input(t *testing.T) {
	expected := ComparableArray{intTest(3), intTest(2), intTest(6), intTest(5), intTest(4), intTest(1)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(5), intTest(4), intTest(6), intTest(1), intTest(3), intTest(2)}

	st.InsertArray(cs)

	ft, err := NewFreeTree(st)
	if err != nil {
		t.Error(err)
	}
	flat := ft.Flatten()

	for i := range flat {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if ft.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(7)) != nil {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(8)) != nil {
		t.Error("unexpected retval")
	}
}

func TestFreeTree_pair_unsorted_input_rebalanced(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(5), intTest(4), intTest(6), intTest(1), intTest(3), intTest(2)}

	st.InsertArray(cs)

	ft, err := NewFreeTree(st.Rebalance())
	if err != nil {
		t.Error(err)
	}
	flat := ft.Flatten()

	for i := range flat {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if ft.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(7)) != nil {
		t.Error("unexpected retval")
	}
	if ft.Ascend(intTest(8)) != nil {
		t.Error("unexpected retval")
	}
}

// -----------------------------------------------------------------------------

type Int int

func (i1 Int) Less(i2 Comparable) bool { return i1 < i2.(Int) }

func Example_simple_usage() {
	// build a new SimpleTree and insert 3 integers in it
	st := NewSimpleTree().Insert(Int(17), Int(66), Int(42))

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
	ft, err := NewFreeTree(st)
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

	// Output:
	// 42
	// <nil>
	// [42 17 66]
	// [17 66 42]
	// 42
	// <nil>
	// [17 66 42]
}
