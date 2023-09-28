package main

import "fmt"

// https://leetcode.com/problems/construct-quad-tree/
func main() {

	x := [][]int{{0, 1}, {1, 0}}
	fmt.Println(x, construct(x))

	x = [][]int{{1, 1}, {1, 1}}
	fmt.Println(x, construct(x))

	x = [][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}}
	fmt.Println(x, construct(x))
}

func construct(grid [][]int) *Node {

	return compactGrid(grid, 0, 0, len(grid))
}

func compactGrid(grid [][]int, i int, j int, len int) *Node {
	node := &Node{Val: grid[i][j] == 1, IsLeaf: false}

	if sameValues(grid, i, j, len) {
		node.IsLeaf = true
		return node
	}

	mid := len / 2
	node.TopLeft = compactGrid(grid, i, j, mid)
	node.TopRight = compactGrid(grid, i, j+mid, mid)
	node.BottomLeft = compactGrid(grid, i+mid, j, mid)
	node.BottomRight = compactGrid(grid, i+mid, j+mid, mid)
	return node
}

func sameValues(grid [][]int, x int, y int, len int) bool {
	sum := 0
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			sum += grid[x+i][y+j]
		}
	}

	return sum == 0 || sum == len*len
}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}
