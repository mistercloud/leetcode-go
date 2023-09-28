package main

import (
	"fmt"
	"sort"
)

func main() {

	x := []int{3, 0, 6, 1, 5}
	fmt.Println(x, hIndex(x))

	x = []int{1, 3, 1}
	fmt.Println(x, hIndex(x))

	x = []int{0, 0, 1}
	fmt.Println(x, hIndex(x))
}

func hIndex(citations []int) int {

	sort.Slice(citations, func(i, j int) bool {
		return citations[i] >= citations[j]
	})

	max := 0
	for i := 1; i <= len(citations); i++ {
		if citations[i-1] >= i {
			max = i
		}
	}

	return max
}
