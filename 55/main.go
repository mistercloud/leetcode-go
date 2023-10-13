package main

// https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums); i++ {
		if i > max {
			return false
		}
		maxFromPoint := i + nums[i]
		if maxFromPoint > max {
			max = maxFromPoint
		}
		if max >= len(nums) {
			return true
		}
	}
	return true
}
