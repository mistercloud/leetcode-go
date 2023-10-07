package main

// https://leetcode.com/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	ans := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = Max(nums[i], sum+nums[i])
		ans = Max(sum, ans)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
