package main

import (
	"fmt"
)

func checkBrackets(s string) bool {
	var stack []rune
	var l bool
	for _, r := range s {
		l = len(stack) == 0
		switch r {

		case '(', '{', '[':
			stack = append(stack, r)
		case ')':

			if l || stack[len(stack)-1] != '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		case '}':

			if l || stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		case ']':

			if l || stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		default:

			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(checkBrackets("{[(}]}"))   // false
	fmt.Println(checkBrackets("{()[][]}")) // true
}
