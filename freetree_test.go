package freetree

import "testing"

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
