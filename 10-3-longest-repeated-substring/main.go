package main

import (
	"bufio"
	"fmt"
	suffixarray "new-way/10-suffix-array"
	"os"
	"strings"
)

func LongestRepeatedSubstring(str string) string {
	sa := suffixarray.New(str)
	suffixes := sa.Arr()
	LCPs := sa.LCP()

	maxLcpIdx := 0
	for i, lcp := range LCPs {
		if lcp > LCPs[maxLcpIdx] {
			maxLcpIdx = i
		}
	}

	longestPrefix := sa.Str()[suffixes[maxLcpIdx]:]
	longestPrefixPrev := sa.Str()[suffixes[maxLcpIdx-1]:]

	lrs := ""
	for i, ch := range longestPrefix {
		if i >= len(longestPrefixPrev) {
			break
		}

		if ch == rune(longestPrefixPrev[i]) {
			lrs += string(ch)
		}
	}

	return lrs
}

func ReadLine(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(line))
}

func main() {
	fmt.Println("<< Longest repeated substring calculator >>")

	reader := bufio.NewReaderSize(os.Stdin, 1024)

	fmt.Println("Enter string:")
	str := ReadLine(reader)

	lrs := LongestRepeatedSubstring(str)

	fmt.Printf("LRS: %s\n", lrs)
}
