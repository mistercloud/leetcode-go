package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-size-subarray-sum/
func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(nums, target, minSubArrayLen(target, nums))

	nums = []int{1, 4, 4}
	target = 1
	fmt.Println(nums, target, minSubArrayLen(target, nums))

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	target = 11
	fmt.Println(nums, target, minSubArrayLen(target, nums))

	nums = []int{}
	target = 1
	fmt.Println(nums, target, minSubArrayLen(target, nums))

	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	target = 3
	fmt.Println(nums, target, minSubArrayLen(target, nums))

}

func minSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	sum := 0
	start, end := 0, 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			minLen = min(minLen, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
