package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := []int{0, 1, 2, 4, 5, 7}
	//fmt.Println(x, summaryRanges(x))

	x = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(x, summaryRanges(x))

	x = []int{}
	fmt.Println(x, summaryRanges(x))

	x = []int{1}
	fmt.Println(x, summaryRanges(x))

	x = []int{1, 2, 3}
	fmt.Println(x, summaryRanges(x))

	x = []int{1, 3, 5}
	fmt.Println(x, summaryRanges(x))
}

// https://leetcode.com/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	var ret []string
	for i := 0; i < len(nums); i++ {
		j := 0
		for i+j < len(nums) {
			if i+j+1 < len(nums) {
				if nums[i]+j+1 == nums[i+j+1] {
					j++
				} else {
					break
				}
			} else {
				break
			}

		}
		ret = append(ret, numRange(nums[i], nums[i+j]))
		i = i + j
	}
	return ret
}

func numRange(a int, b int) string {
	if a == b {
		return strconv.Itoa(a)
	}

	return strconv.Itoa(a) + `->` + strconv.Itoa(b)
}
