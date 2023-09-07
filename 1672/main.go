package main

import "fmt"

func main() {
	x := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(x, ` - `, maximumWealth(x))
	x = [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(x, ` - `, maximumWealth(x))
	x = [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(x, ` - `, maximumWealth(x))
}

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := range accounts {
		sum := 0
		for j := range accounts[i] {
			sum += accounts[i][j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
