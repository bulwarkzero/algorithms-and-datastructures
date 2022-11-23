package main

import (
	"bufio"
	"fmt"
	suffixarray "new-way/10-suffix-array"
	"os"
	"strconv"
	"strings"
)

// TODO: dynamically generate sentinels
// Unique sentinels used to concatenate strings
// A sentinel must be lexicographically less that any characters in strings
var sentinels = []rune{0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a}

func concatenateStrings(strings []string) string {
	concatenated := ""

	for i := 0; i < len(strings); i++ {
		concatenated += strings[i] + string(sentinels[i])
	}

	return concatenated
}

func isSentinel(ch rune) bool {
	return ch >= sentinels[0] && ch <= sentinels[len(sentinels)-1]
}

func sentinelsInWindow(str []rune, suffixSentinels []rune, windowStart, windowEnd int) int {
	seenSentinels := make(map[rune]struct{})

	for i := windowStart; i <= windowEnd; i++ {
		seenSentinels[suffixSentinels[i]] = struct{}{}
	}

	return len(seenSentinels)
}

func arrMin(arr []int) int {
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}

func lcp(str []rune, suffixes []int, LCPs []int, start, end int) []rune {
	minLCP := arrMin(LCPs[start+1 : end+1])
	if minLCP == 0 {
		return nil
	}

	return str[suffixes[start] : suffixes[start]+minLCP]
}

func LongestCommonSubstring(strings []string, k int) string {
	concatenatedString := concatenateStrings(strings)
	concatenatedStringRune := []rune(concatenatedString)
	stringsLen := len(strings)

	sa := suffixarray.New(concatenatedString)
	suffixes := sa.Arr()
	LCPs := sa.LCP()

	strSentinels := make([]rune, len(concatenatedStringRune))
	currentSentinel := sentinels[0]
	for i := len(concatenatedStringRune) - 1; i >= 0; i-- {
		chr := concatenatedStringRune[i]

		if isSentinel(chr) {
			currentSentinel = chr
		}

		strSentinels[i] = currentSentinel
	}

	suffixSentinels := make([]rune, len(strSentinels))
	for i, suffixIdx := range suffixes {
		suffixSentinels[i] = strSentinels[suffixIdx]
	}

	// first suffixes are sentinels, we omit those
	windowStart := stringsLen
	windowEnd := stringsLen + k - 1
	var lcs []rune

	for windowEnd < len(suffixes) {
		numSentinelsInWindow := sentinelsInWindow(concatenatedStringRune, suffixSentinels, windowStart, windowEnd)
		if numSentinelsInWindow < k {
			windowEnd += 1
			continue
		}

		lcp := lcp(concatenatedStringRune, suffixes, LCPs, windowStart, windowEnd)
		if len(lcp) >= len(lcs) {
			lcs = lcp
		}

		windowStart += 1
	}

	return string(lcs)
}

func ReadLine(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(line))
}

func main() {
	fmt.Println("<< Longest common substring calculator >>")

	reader := bufio.NewReaderSize(os.Stdin, 1024)

	fmt.Println("Enter number of strings:")
	tmp := ReadLine(reader)

	numStrings, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}

	strings := make([]string, numStrings)

	for i := 0; i < numStrings; i++ {
		fmt.Printf("Enter %dst string:\n", i+1)
		strings[i] = ReadLine(reader)
	}

	fmt.Println("Enter k (number of strings that must have common substring):")
	tmp = ReadLine(reader)

	k, err := strconv.Atoi(tmp)
	if err != nil {
		panic(err)
	}

	if k < 2 {
		fmt.Printf("k must be at least 2, got %d\n", k)
		return
	}

	if k > len(strings) {
		fmt.Printf("k is greater than number of provided strings %d\n", k)
		return
	}

	lcs := LongestCommonSubstring(strings, k)

	fmt.Printf("LCS: %s\n", lcs)
}
