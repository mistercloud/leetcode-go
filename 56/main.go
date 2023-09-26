package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/merge-intervals/
func main() {
	x := [][]int{{2, 6}, {1, 3}, {15, 18}, {8, 10}}
	fmt.Println(x, merge(x))

	x = [][]int{{1, 4}, {4, 5}}
	fmt.Println(x, merge(x))

	x = [][]int{{1, 4}}
	fmt.Println(x, merge(x))

	x = [][]int{{1, 4}, {1, 3}}
	fmt.Println(x, merge(x))

	x = [][]int{{1, 4}, {2, 3}}
	fmt.Println(x, merge(x))
}

func merge(intervals [][]int) [][]int {

	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var merged [][]int
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else if merged[len(merged)-1][1] < interval[1] {
			merged[len(merged)-1][1] = interval[1]
		}
	}
	return merged
}
