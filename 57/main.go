package main

// https://leetcode.com/problems/insert-interval/
func insert(intervals [][]int, newInterval []int) [][]int {
	var ret [][]int

	for i := 0; i < len(intervals); i++ {
		if newInterval[1] < intervals[i][0] {
			ret = append(ret, newInterval)
			ret = append(ret, intervals[i:]...)

			return ret
		} else if newInterval[0] > intervals[i][1] {
			ret = append(ret, intervals[i])
		} else {
			newInterval = []int{min(newInterval[0], intervals[i][0]), max(newInterval[1], intervals[i][1])}
		}
	}

	return append(ret, newInterval)
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if y > x {
		return y
	}

	return x
}
