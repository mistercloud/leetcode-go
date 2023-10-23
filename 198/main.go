package _198

// https://leetcode.com/problems/house-robber/
func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i] = nums[i]
		if i > 2 {
			dp[i] += max(dp[i-2], dp[i-3])
		} else {
			dp[i] += dp[i-2]
		}

	}

	return max(dp[len(dp)-1], dp[len(dp)-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
