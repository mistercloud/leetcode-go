package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := `00000000000000000000000000001011`
	xs, _ := strconv.ParseUint(x, 2, 32)
	xs32 := uint32(xs)
	fmt.Println(x, hammingWeight(xs32))

	x = `00000000000000000000000010000000`
	xs, _ = strconv.ParseUint(x, 2, 32)
	xs32 = uint32(xs)
	fmt.Println(x, hammingWeight(xs32))

	x = `11111111111111111111111111111101`
	xs, _ = strconv.ParseUint(x, 2, 32)
	xs32 = uint32(xs)
	fmt.Println(x, hammingWeight(xs32))
}

func hammingWeight(num uint32) int {

	if num == 0 {
		return 0
	}
	return int(num&1) + hammingWeight(num>>1)
}
