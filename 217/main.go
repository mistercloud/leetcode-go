package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 1}
	fmt.Println(arr, ` - `, containsDuplicate(arr))
	arr = []int{1, 2, 3, 4}
	fmt.Println(arr, ` - `, containsDuplicate(arr))
	arr = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(arr, ` - `, containsDuplicate(arr))
	arr = []int{1}
	fmt.Println(arr, ` - `, containsDuplicate(arr))

	arr = []int{1, 2, 4, 2}
	fmt.Println(arr, ` - `, containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	//решение O(n + log n) + O (n)
	/*
		sort.IntSlice(nums).Sort()
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] == nums[i+1] {
				return true
			}
		}
		return false
	*/
	//решение O(n) но с неизвестным размером hmap
	hmap := make(map[int]int, len(nums))
	for _, val := range nums {
		hmap[val]++
		if hmap[val] > 1 {
			return true
		}
	}

	return false
}
