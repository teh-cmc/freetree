package freetree

import (
	"log"
	"testing"
)

// -----------------------------------------------------------------------------

type intTest int

func (i1 intTest) Less(i2 Comparable) bool { return i1 < i2.(intTest) }

func TestSimpleTree_pair_sorted_input(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(1), intTest(2), intTest(3), intTest(4), intTest(5), intTest(6)}

	st.InsertArray(cs)
	flat := st.Flatten()

	for i := range expected {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if st.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(7)) != nil {
		log.Println(st.Ascend(intTest(7)).(intTest))
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(8)) != nil {
		log.Println(st.Ascend(intTest(8)).(intTest))
		t.Error("unexpected retval")
	}
}

func TestSimpleTree_pair_sorted_input_rebalanced(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(1), intTest(2), intTest(3), intTest(4), intTest(5), intTest(6)}

	st.InsertArray(cs)
	flat := st.Flatten()

	for i := range expected {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if st.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(7)) != nil {
		log.Println(st.Ascend(intTest(7)).(intTest))
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(8)) != nil {
		log.Println(st.Ascend(intTest(8)).(intTest))
		t.Error("unexpected retval")
	}
}

func TestSimpleTree_pair_unsorted_input(t *testing.T) {
	expected := ComparableArray{intTest(3), intTest(2), intTest(6), intTest(5), intTest(4), intTest(1)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(5), intTest(4), intTest(6), intTest(1), intTest(3), intTest(2)}

	st.InsertArray(cs)
	flat := st.Flatten()

	for i := range expected {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if st.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(7)) != nil {
		log.Println(st.Ascend(intTest(7)).(intTest))
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(8)) != nil {
		log.Println(st.Ascend(intTest(8)).(intTest))
		t.Error("unexpected retval")
	}
}

func TestSimpleTree_pair_unsorted_input_rebalanced(t *testing.T) {
	expected := ComparableArray{intTest(1), intTest(3), intTest(2), intTest(5), intTest(6), intTest(4)}

	st := NewSimpleTree()
	cs := ComparableArray{intTest(5), intTest(4), intTest(6), intTest(1), intTest(3), intTest(2)}

	st.InsertArray(cs)
	st.Rebalance()
	flat := st.Flatten()

	for i := range expected {
		if flat[i] != expected[i] {
			t.Error("expected != flat")
			t.FailNow()
		}
	}

	if st.Ascend(intTest(1)).(intTest) != intTest(1) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(2)).(intTest) != intTest(2) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(5)).(intTest) != intTest(5) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(6)).(intTest) != intTest(6) {
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(7)) != nil {
		log.Println(st.Ascend(intTest(7)).(intTest))
		t.Error("unexpected retval")
	}
	if st.Ascend(intTest(8)) != nil {
		log.Println(st.Ascend(intTest(8)).(intTest))
		t.Error("unexpected retval")
	}
}
