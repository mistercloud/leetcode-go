package main

import "fmt"

// https://leetcode.com/problems/search-insert-position/
func main() {

	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(nums, target, searchInsert(nums, target), 2)

	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(nums, target, searchInsert(nums, target), 1)

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(nums, target, searchInsert(nums, target), 4)

	nums = []int{1, 3, 5, 6}
	target = -1
	fmt.Println(nums, target, searchInsert(nums, target), 0)

	nums = []int{2}
	target = 1
	fmt.Println(nums, target, searchInsert(nums, target), 0)

	nums = []int{2}
	target = 3
	fmt.Println(nums, target, searchInsert(nums, target), 1)
}

func searchInsert(nums []int, target int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return searchInsert(nums[:mid], target)
	}
	return mid + searchInsert(nums[mid:], target)
}
