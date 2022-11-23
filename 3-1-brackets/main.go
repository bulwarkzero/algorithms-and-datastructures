package main

import (
	"fmt"
	stack "new-way/3-stack"
)

func isClosedBracket(bracket byte) bool {
	return bracket == ']' || bracket == ')' || bracket == '}'
}

func isOpenedBracket(bracket byte) bool {
	return bracket == '[' || bracket == '(' || bracket == '{'
}

func isBracketsMatch(openBracket, closeBracket byte) bool {
	switch openBracket {
	case '[':
		return closeBracket == ']'
	case '{':
		return closeBracket == '}'
	case '(':
		return closeBracket == ')'
	}

	return false
}

func isBracketsValid(brackets string) bool {
	stack := stack.NewArrayStack()

	for _, bracket := range []byte(brackets) {
		if isOpenedBracket(bracket) {
			stack.Push(int(bracket))
		} else {
			if stack.Len() < 1 {
				return false
			}

			openBracket := byte(stack.Pop())

			if !isBracketsMatch(openBracket, bracket) {
				return false
			}
		}
	}

	return true
}

func main() {
	bracketsStr := "[{{}}][()[[[]]]]"

	fmt.Printf("Is Brackets: %s are valid: %v\r\n", bracketsStr, isBracketsValid(bracketsStr))
}
