package main

// https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	matrixLen := len(matrix) * len(matrix[0])
	end := matrixLen - 1
	start := 0
	for start <= end {
		mid := (start + end) / 2
		//переводим число в координаты
		j := mid % len(matrix[0])
		i := mid / len(matrix[0])
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}
