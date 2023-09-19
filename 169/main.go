package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3}
	expected := 3
	fmt.Println(nums, majorityElement(nums), expected, majorityElement(nums) == expected)

	nums = []int{3, 2, 3}
	expected = 3
	fmt.Println(nums, majorityElement(nums), expected, majorityElement(nums) == expected)

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	expected = 2
	fmt.Println(nums, majorityElement(nums), expected, majorityElement(nums) == expected)
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
