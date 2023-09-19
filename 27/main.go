package main

import "fmt"

func main() {

	x := []int{3, 2, 2, 3}
	num := 3
	fmt.Println(x, ` - `, num)
	fmt.Println(removeElement(x, num), ` `, x)

	x = []int{2, 2, 3}
	num = 2
	fmt.Println(x, ` - `, num)
	fmt.Println(removeElement(x, num), ` `, x)

	x = []int{2, 2, 2}
	num = 2
	fmt.Println(x, ` - `, num)
	fmt.Println(removeElement(x, num), ` `, x)

	x = []int{3, 3, 3}
	num = 2
	fmt.Println(x, ` - `, num)
	fmt.Println(removeElement(x, num), ` `, x)

	x = []int{0, 4, 4, 0, 4, 4, 4, 0, 2}
	num = 4
	fmt.Println(x, ` - `, num)
	fmt.Println(removeElement(x, num), ` `, x)
}

// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
