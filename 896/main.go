package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{6, 5, 4, 4}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{1, 2, 2, 3}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{1, 3, 2}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{1}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{1, 2}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{2, 1}
	fmt.Println(arr, ` - `, isMonotonic(arr))
	arr = []int{1, 1, 2}
	fmt.Println(arr, ` - `, isMonotonic(arr))
}

// https://leetcode.com/problems/monotonic-array/description/
func isMonotonic(nums []int) bool {
	asc, desc := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			asc = false
		} else if nums[i] < nums[i+1] {
			desc = false
		}

		if !asc && !desc {
			return false
		}
	}

	return asc || desc
}
