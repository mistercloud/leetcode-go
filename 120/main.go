package main

import (
	"fmt"
)

// https://leetcode.com/problems/triangle/
func main() {

	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}

	fmt.Println(minimumTotal(triangle))

	triangle = [][]int{
		{-10},
	}

	fmt.Println(minimumTotal(triangle))

	triangle = [][]int{
		{-10},
		{-10, 2},
	}

	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
