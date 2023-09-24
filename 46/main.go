package main

import (
	"fmt"
)

// https://leetcode.com/problems/permutations/
func main() {

	nums := []int{1}
	//fmt.Println(nums, permute(nums), len(permute(nums)))

	nums = []int{1, 2}
	fmt.Println(nums, permute(nums), len(permute(nums)))

	nums = []int{1, 2, 3}
	fmt.Println(nums, permute(nums), len(permute(nums)))

	nums = []int{1, 2, 3, 4}
	fmt.Println(nums, permute(nums), len(permute(nums)))
}

func permute(nums []int) [][]int {
	var ret [][]int
	permitSubArray(nums, 0, &ret)
	return ret
}

func permitSubArray(nums []int, start int, ret *[][]int) {

	if start == len(nums)-1 {
		tst := make([]int, len(nums))
		copy(tst, nums)
		*ret = append(*ret, tst)
		return
	}
	for j := start; j < len(nums); j++ {
		nums[start], nums[j] = nums[j], nums[start]
		permitSubArray(nums, start+1, ret)
		nums[start], nums[j] = nums[j], nums[start]
	}
}
