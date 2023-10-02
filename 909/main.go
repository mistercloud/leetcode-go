package main

import (
	"fmt"
)

func main() {

	board := [][]int{
		{-1, -1},
		{-1, 3},
	}
	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1},
	}

	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{5, 5, 5, 5, 5, 5},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}

	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, -1},
		{-1, 9, 8},
		{-1, 8, 9},
	}

	fmt.Println(snakesAndLadders(board))

	board = [][]int{
		{-1, -1, 2, -1},
		{14, 2, 12, 3},
		{4, 9, 1, 11},
		{-1, 2, 1, 16},
	}

	fmt.Println(snakesAndLadders(board))

}

func snakesAndLadders(board [][]int) int {
	target := len(board) * len(board)
	if target == 4 {
		//нужен одина шаг
		return 1
	}
	dp := make([]int, target+1)
	moves := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = target
	}

	//переводим 2d в 1d доску
	pos, rowCnt := 1, 0
	for i := len(board) - 1; i >= 0; i-- {
		if rowCnt == 0 {
			for rowCnt < len(board) {
				moves[pos] = board[i][rowCnt]
				pos++
				rowCnt++
			}
		} else {
			for rowCnt > 0 {
				rowCnt--
				moves[pos] = board[i][rowCnt]
				pos++

			}
		}
	}

	//BFS queue
	q := make([]int, 0)
	q = append(q, 1)
	dp[1] = 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		//проходим по следующим клеткам
		for j := i + 1; j <= Min(i+6, target); j++ {
			//если на клетке переход
			if moves[j] > 0 {
				if dp[moves[j]] > dp[i]+1 {
					q = append(q, moves[j])
					dp[moves[j]] = dp[i] + 1
				}
			} else {
				if dp[j] > dp[i]+1 {
					q = append(q, j)
					dp[j] = dp[i] + 1
				}
			}

		}
	}

	if dp[target] == target {
		return -1
	}
	return dp[target]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
