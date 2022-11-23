package fenwicktree

import "testing"

func TestConstruct(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	tree := New(arr)

	if tree.arr[2] != 3 {
		t.Fatalf("construction error, index 2 expected to be 3, got %d", tree.arr[2])
	}

	if tree.arr[3] != 10 {
		t.Fatalf("construction error, index 3 expected to be 10, got %d", tree.arr[3])
	}

	if tree.arr[5] != 11 {
		t.Fatalf("construction error, index 5 expected to be 11, got %d", tree.arr[5])
	}
}

func TestPrefixSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	tree := New(arr)

	if ps := tree.PrefixSum(1); ps != 1 {
		t.Fatalf("prefixSum(1) expected to be 1, got %d", ps)
	}

	if ps := tree.PrefixSum(2); ps != 3 {
		t.Fatalf("prefixSum(2) expected to be 3, got %d", ps)
	}

	if ps := tree.PrefixSum(4); ps != 10 {
		t.Fatalf("prefixSum(4) expected to be 10, got %d", ps)
	}

	if ps := tree.PrefixSum(6); ps != 21 {
		t.Fatalf("prefixSum(6) expected to be 21, got %d", ps)
	}
}

func TestQuery(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	tree := New(arr)

	if sum := tree.Query(3, 6); sum != 18 {
		t.Fatalf("query(3, 6) expected to be 18, got %d", sum)
	}

	if sum := tree.Query(4, 6); sum != 15 {
		t.Fatalf("query(4, 6) expected to be 15, got %d", sum)
	}

	if sum := tree.Query(2, 6); sum != 20 {
		t.Fatalf("query(4, 6) expected to be 20, got %d", sum)
	}
}

func TestUpdate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	tree := New(arr)

	tree.Add(4, 5)

	if sum := tree.PrefixSum(4); sum != 15 {
		t.Fatalf("prefixSum(4) after update expected to be 15, got %d", sum)
	}
}
