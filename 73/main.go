package main

import "fmt"

// https://leetcode.com/problems/set-matrix-zeroes/
func main() {

	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{1, 0},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {

	nulRow := false
	nulCol := false
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			nulCol = true
			break
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			nulRow = true
			break
		}
	}

	//проходим по всем внутренним ячейкам
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	//идем по первому ряду выставляем 0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if nulRow {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if nulCol {
		for j := 0; j < len(matrix); j++ {
			matrix[j][0] = 0
		}
	}
}
