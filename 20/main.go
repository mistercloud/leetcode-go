package main

import (
	"fmt"
)

// https://leetcode.com/problems/valid-parentheses/
func main() {
	s := `()[]{}`
	fmt.Println(s, isValid(s))

	s = `((`
	fmt.Println(s, isValid(s))

	s = `([{}()])`
	fmt.Println(s, isValid(s))
	s = `([{}(()])`
	fmt.Println(s, isValid(s))

	s = `()[]{}`
	fmt.Println(s, isValid(s))

	s = `(]`
	fmt.Println(s, isValid(s))
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := NewStack()
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack.Push(ch)
		} else {
			last := stack.Pop()
			if ch == ')' && last != '(' {
				return false
			}
			if ch == '}' && last != '{' {
				return false
			}
			if ch == ']' && last != '[' {
				return false
			}
		}
	}
	return stack.isEmpty()
}

type Stack struct {
	values []rune
	top    int
}

func NewStack() *Stack {
	return &Stack{
		values: make([]rune, 0),
		top:    -1,
	}
}

func (stack *Stack) Push(val rune) {
	stack.top++
	if stack.top < len(stack.values) {
		stack.values[stack.top] = val
	} else {
		stack.values = append(stack.values, val)
	}

}

func (stack *Stack) Pop() rune {
	if stack.isEmpty() {
		return 0
	}
	value := stack.values[stack.top]
	stack.top--
	return value
}
func (stack *Stack) isEmpty() bool {
	return stack.top == -1
}
