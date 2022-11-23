package suffixarray

import (
	"sort"
	"strings"
	"unicode/utf8"
)

type SuffixArray struct {
	str string
	arr []int
}

func (sa *SuffixArray) Arr() []int {
	return sa.arr
}

func (sa *SuffixArray) Str() string {
	return sa.str
}

func (sa *SuffixArray) LCP() []int {
	lcp := make([]int, len(sa.arr))
	inv := make([]int, len(sa.arr))
	lcpLen := len(lcp)

	for i, suffix := range sa.arr {
		inv[suffix] = i
	}

	suffixLen := 0
	for i := 0; i < lcpLen; i++ {
		if inv[i] == 0 {
			continue
		}

		j := sa.arr[inv[i]-1]

		for (j+suffixLen < lcpLen) && (i+suffixLen < lcpLen) && (sa.str[j+suffixLen] == sa.str[i+suffixLen]) {
			suffixLen++
		}

		lcp[inv[i]] = suffixLen

		if suffixLen > 0 {
			suffixLen -= 1
		}

	}

	return lcp
}

func (sa *SuffixArray) Find(pat string) int {
	patLen := utf8.RuneCountInString(pat)
	strLen := len(sa.arr)
	lo := 0
	hi := strLen

	for lo <= hi {
		mid := (lo + (hi - 1)) / 2
		midSuffix := sa.arr[mid]
		searchEndIdx := midSuffix + patLen
		if searchEndIdx > strLen {
			searchEndIdx = strLen
		}

		res := strings.Compare(pat, sa.str[midSuffix:searchEndIdx])

		if res == 0 {
			return sa.arr[mid]
		}

		if res < 0 {
			hi = mid
		} else {
			lo = mid
		}
	}

	return -1
}

// TODO: implement O(n*log(n)) algorithm
func New(s string) *SuffixArray {
	strLen := utf8.RuneCountInString(s)
	suffixes := make([]string, strLen)
	suffixArray := make([]int, strLen)

	for i := 0; i < strLen; i++ {
		suffixes[i] = s[i:]
	}

	sort.Strings(suffixes)

	for i, suffix := range suffixes {
		suffixArray[i] = strLen - utf8.RuneCountInString(suffix)
	}

	return &SuffixArray{str: s, arr: suffixArray}
}
