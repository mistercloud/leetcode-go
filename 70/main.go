package main

import "fmt"

func main() {
	n := 1
	fmt.Println(n, climbStairs(n))

	n = 2
	fmt.Println(n, climbStairs(n))

	n = 3
	fmt.Println(n, climbStairs(n))

	n = 45
	fmt.Println(n, climbStairs(n))
}

// https://leetcode.com/problems/climbing-stairs
func climbStairs(n int) int {

	if n <= 0 {
		return 0
	}
	ways := [2]int{0, 1}
	for i := 1; i <= n; i++ {
		ways[0], ways[1] = ways[1], ways[0]+ways[1]
	}

	return ways[1]

}
