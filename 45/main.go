package _45

import "math"

// https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) int {

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if dp[i+j] == math.MaxInt {
				dp[i+j] = min(dp[i+j], dp[i]+1)
				if i+j == len(nums)-1 {
					return dp[len(nums)-1]
				}
			}

		}
	}

	return dp[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
