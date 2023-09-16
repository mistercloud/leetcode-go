package main

import "fmt"

func main() {
	x := 2
	fmt.Println(x, ` - `, isThree(x), ` - `, false)
	x = 49
	fmt.Println(x, ` - `, isThree(x), ` - `, true)
	x = 3
	fmt.Println(x, ` - `, isThree(x), ` - `, false)
	x = 100
	fmt.Println(x, ` - `, isThree(x), ` - `, false)
	x = 25
	fmt.Println(x, ` - `, isThree(x), ` - `, true)
	x = 15
	fmt.Println(x, ` - `, isThree(x), ` - `, false)
	x = 7
	fmt.Println(x, ` - `, isThree(x), ` - `, false)
	x = 17
	fmt.Println(x, ` - `, isThree(x), ` - `, false)

}

// https://leetcode.com/problems/three-divisors/submissions/1051187108/
func isThree(n int) bool {
	if n < 4 {
		return false
	}
	cnt := 2
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			cnt++
		}

		if cnt > 3 {
			return false
		}
	}
	return cnt == 3
}
