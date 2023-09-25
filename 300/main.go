package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
func main() {
	x := []int{1}
	fmt.Println(x, lengthOfLIS(x))

	x = []int{1, 0}
	fmt.Println(x, lengthOfLIS(x))

	x = []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(x, lengthOfLIS(x))

	x = []int{0, 1, 0, 3, 2, 3}
	fmt.Println(x, lengthOfLIS(x))

	x = []int{7, 7, 7, 7, 7, 7, 7}
	fmt.Println(x, lengthOfLIS(x))

	x = []int{101, 102, 103, 10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(x, lengthOfLISBinary(x))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}

	max := 1
	for _, num := range dp {
		max = Max(max, num)
	}

	return max
}

// решение через DP + бинарный поиск
func lengthOfLISBinary(nums []int) int {
	var l []int
	for i := 0; i < len(nums); i++ {
		pos := sort.Search(len(l), func(j int) bool {
			return l[j] >= nums[i]
		})
		if pos >= len(l) {
			l = append(l, nums[i])
		} else {
			l[pos] = nums[i]
		}

		fmt.Println(i, l)
	}
	return len(l)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
