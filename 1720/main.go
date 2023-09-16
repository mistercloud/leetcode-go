package main

import "fmt"

func main() {

	x := []int{1, 2, 3}
	fmt.Println(x, ` - `, decode(x, 1), ` expected `, []int{1, 0, 2, 1})
	x = []int{6, 2, 7, 3}
	fmt.Println(x, ` - `, decode(x, 4), ` expected `, []int{4, 2, 0, 7, 4})
}

// https://leetcode.com/problems/decode-xored-array/submissions/1051156050/
func decode(encoded []int, first int) []int {
	var decoded = make([]int, len(encoded)+1)
	decoded[0] = first
	for i := 1; i < len(decoded); i++ {
		decoded[i] = decoded[i-1] ^ encoded[i-1]
	}
	return decoded

}
