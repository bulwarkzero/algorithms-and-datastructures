package main

import (
	"bufio"
	"fmt"
	suffixarray "new-way/10-suffix-array"
	"os"
	"strings"
	"unicode/utf8"
)

func numUniqueSubStrings(str string) int {
	strLen := utf8.RuneCountInString(str)

	sa := suffixarray.New(str)
	lcp := sa.LCP()
	totalSubStrings := strLen * (strLen + 1) / 2

	lcpSum := 0
	for i := 1; i < strLen; i++ {
		lcpSum += lcp[i]
	}

	return totalSubStrings - lcpSum
}

func main() {
	fmt.Println("<< Unique substrings calculator >>")
	fmt.Println("Enter your string:")

	reader := bufio.NewReader(os.Stdin)

	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	str := strings.TrimSpace(string(line))

	uniqueSubstringsCount := numUniqueSubStrings(str)

	fmt.Println(uniqueSubstringsCount)
}
