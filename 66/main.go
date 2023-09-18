package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Print(x, ` `)
	fmt.Println(plusOne(x))

	x = []int{4, 3, 2, 1}
	fmt.Print(x, ` `)
	fmt.Println(plusOne(x))

	x = []int{9}
	fmt.Print(x, ` `)
	fmt.Println(plusOne(x))

	x = []int{9, 9}
	fmt.Print(x, ` `)
	fmt.Println(plusOne(x))
}

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 10 && i != 0 {
			digits[i-1]++
			digits[i] = 0
		} else {
			break
		}
	}
	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}
