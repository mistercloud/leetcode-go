package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[left]
	}

	for left <= right {
		mid := (left + right) / 2
		if right-left == 1 {
			return min(nums[left], nums[right])
		}
		if nums[mid] > nums[right] {
			left = mid
		} else {
			right = mid
		}
	}
	return nums[left]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
