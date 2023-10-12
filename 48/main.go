package main

// https://leetcode.com/problems/rotate-image/
func rotate(matrix [][]int) {

	n := len(matrix)
	//транспонируем по диагонали
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//отзеркаливаем по вертикали
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}

}
