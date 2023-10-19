package _38

// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {

	dpleft := make([]int, len(nums))
	dpleft[0] = 1
	for i := 1; i < len(nums); i++ {
		dpleft[i] = dpleft[i-1] * nums[i-1]
	}

	dpright := make([]int, len(nums))
	dpright[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		dpright[i] = dpright[i+1] * nums[i+1]
	}

	ret := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = dpleft[i] * dpright[i]
	}

	return ret
}
