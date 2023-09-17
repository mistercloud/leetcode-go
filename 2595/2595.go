package main

import (
	"fmt"
)

func main() {

	x := 17
	fmt.Println(x, ` - `, evenOddBit(x), ` = `, []int{2, 0})

	x = 2
	fmt.Println(x, ` - `, evenOddBit(x), ` = `, []int{0, 1})
	x = 3
	fmt.Println(x, ` - `, evenOddBit(x), ` = `, []int{1, 1})
}

// https://leetcode.com/problems/number-of-even-and-odd-bits/submissions/1051621911/
func evenOddBit(n int) []int {
	evenOdd := make([]int, 2)
	i := 0
	for n > 0 {
		if n%2 == 1 {
			evenOdd[i]++
		}

		n = n / 2
		if i == 0 {
			i = 1
		} else {
			i = 0
		}
	}
	return evenOdd
}

//https://leetcode.com/problems/number-of-even-and-odd-bits/submissions/1051621911/
/*func evenOddBit(n int) []int {

	evenOdd := make([]int, 2)
	bin := strconv.FormatInt(int64(n), 2)
	for i, char := range bin {
		if char == 49 {
			evenOdd[(len(bin)-1-i)%2]++
		}
	}
	return evenOdd
}*/
