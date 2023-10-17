package main

import "math/bits"

// https://leetcode.com/problems/reverse-bits/
func reverseBits(num uint32) uint32 {

	return bits.Reverse32(num)
}
