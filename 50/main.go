package main

import "math"

// https://leetcode.com/problems/powx-n/
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	return math.Pow(x, float64(n))
}
