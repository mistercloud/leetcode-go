package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nums)
	k := removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
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

	nums = []int{1, 2, 3}
	fmt.Println(nums)
	k = removeDuplicates(nums)
	fmt.Println(nums, k)
}

func removeDuplicates(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}
	index := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[index-2] {
			nums[index] = nums[i]
			index++
		}
	}
	return index

	//предыдущее решение , более запутанное
	i := 0
	j := 1
	foundCopy := false
	for i < len(nums) {
		if j >= len(nums) {
			break
		}
		if i == j {
			j++
		}

		for j < len(nums) {
			if nums[j] > nums[i] {
				nums[i+1], nums[j] = nums[j], nums[i+1]
				foundCopy = false
				break
			} else if nums[j] == nums[i] && !foundCopy {
				nums[i+1], nums[j] = nums[j], nums[i+1]
				foundCopy = true
				break
			}
			if j == len(nums)-1 {
				return i + 1
			}
			j++
		}
		i++
	}

	return i
}
