package main

import "fmt"

func main() {
	a := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println(a)
	gameOfLife(a)
	fmt.Println(a)

	a = [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0}}

	fmt.Println(a)
	gameOfLife(a)
	fmt.Println(a)
}

func gameOfLife(board [][]int) {
	nextState := make([][]int, len(board))
	for i := range nextState {
		nextState[i] = make([]int, len(board[0]))
	}

	//ran := []int{-1, 0, 1}
	for i := 0; i < len(nextState); i++ {
		for j := 0; j < len(nextState[0]); j++ {
			liveCells := 0
			for k := i - 1; k <= i+1 && k < len(nextState); k++ {
				if k >= 0 {
					for n := j - 1; n <= j+1 && n < len(nextState[0]); n++ {
						if n >= 0 && board[k][n] == 1 {
							if !(k == i && n == j) {
								liveCells++
							}
						}
					}
				}
			}

			fmt.Println(i, j, liveCells)
			if board[i][j] == 1 {
				if liveCells < 2 {
					//dies as if caused by under-population.
					nextState[i][j] = 0
				} else if liveCells > 3 {
					// dies, as if by over-population
					nextState[i][j] = 0
				} else {
					//lives on to the next generation.
					nextState[i][j] = 1
				}
			} else if liveCells == 3 {
				// live cell, as if by reproduction
				nextState[i][j] = 1
			}
		}

	}

	for i := 0; i < len(nextState); i++ {
		for j := 0; j < len(nextState[0]); j++ {
			board[i][j] = nextState[i][j]
		}

	}
}
