package main

import "testing"

func Test_maxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			`Test 1`,
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			`Test 2`,
			[]int{1},
			1,
		},
		{
			`Test 3`,
			[]int{5, 4, -1, 7, 8},
			23,
		},
		{
			`Test 4`,
			[]int{-1},
			-1,
		},
		{
			`Test 5`,
			[]int{-5, -6, -7, -1},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
