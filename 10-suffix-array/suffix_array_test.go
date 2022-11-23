package suffixarray

import (
	"testing"
)

func TestNew(t *testing.T) {
	s := "ABCA"

	suffixArray := New(s)
	arr := suffixArray.Arr()

	desiredSuffixArray := []int{3, 0, 1, 2}

	for i, desiredSuffix := range desiredSuffixArray {
		if arr[i] != desiredSuffix {
			t.Fatalf("Expected suffix on index %d to be %d, got %d", i, desiredSuffix, arr[i])
		}
	}
}

func TestLCP(t *testing.T) {
	s := "ABCA"

	suffixArray := New(s)
	lcp := suffixArray.LCP()

	desiredLCP := []int{0, 1, 0, 0}

	for i, dLCP := range desiredLCP {
		if lcp[i] != dLCP {
			t.Fatalf("Expected lcp on index %d to be %d, got %d", i, dLCP, lcp[i])
		}
	}
}

func TestFind(t *testing.T) {
	s := "HELLOFROMHERE"

	suffixArray := New(s)

	if idx := suffixArray.Find("FROM"); idx != 5 {
		t.Fatalf("Expected to find pattern in index 5, got %d", idx)
	}
}
