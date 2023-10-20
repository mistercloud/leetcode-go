package _189

// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {

	//реверт всего массива
	low, high := 0, len(nums)-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}

	k = k % len(nums)

	//реверт левой части до k
	low, high = 0, k-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}

	//реверт правой части от k
	low, high = k, len(nums)-1
	for low < high {
		nums[low], nums[high] = nums[high], nums[low]
		low += 1
		high -= 1
	}

}
