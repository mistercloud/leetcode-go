package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	k := removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1, 2, 2}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1, 1, 1}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1, 1}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1, 2}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {

	if len(nums) == 1 {
		return 1
	}

	i := 0
	j := 1
	for i < len(nums) {
		//ищем следующее уникальное число
		if j >= len(nums) {
			break
		}
		for j < len(nums) {
			if nums[j] != nums[i] {
				nums[i+1] = nums[j]
				break
			}
			j++
		}

		i++
	}
	return i
}
