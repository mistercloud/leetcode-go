package _4

import (
	"sort"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	left := sort.SearchInts(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	//число есть в массиве, двигаем окно чтобы выяснить на каких позициях
	right := left
	for {
		left--
		if left < 0 || nums[left] != target {
			left++
			break
		}
	}
	for {
		right++
		if right >= len(nums) || nums[right] != target {
			right--
			break
		}
	}

	return []int{left, right}
}
