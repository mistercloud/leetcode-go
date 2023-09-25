package main

import "fmt"

// https://leetcode.com/problems/maximal-square/
func main() {
	x := [][]byte{{'0'}}
	fmt.Println(x, maximalSquare(x))

	x = [][]byte{{'1'}}
	fmt.Println(x, maximalSquare(x))

	x = [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	fmt.Println(x, maximalSquare(x))

	x = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '0'},
	}
	fmt.Println(x, maximalSquare(x))

	x = [][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '0'},
		{'1', '1', '1', '1', '1'},
	}
	fmt.Println(x, maximalSquare(x))
}

func maximalSquare(matrix [][]byte) int {
	max := 0
	if matrix[0][0] == '1' {
		max = 1
	}
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] += Min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
				}

				if dp[i][j] > max {
					max = dp[i][j]
				}
			}
		}
	}
	return max * max
}

func Min(x int, y int, z int) int {
	if x <= y {
		y = x
	}
	if y <= z {
		return y
	}
	return z

}
