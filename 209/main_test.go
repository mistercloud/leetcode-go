package main

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			`Test 1`,
			args{target: 7, nums: []int{2, 3, 1, 2, 4, 3}},
			2,
		},
		{
			`Test 2`,
			args{target: 1, nums: []int{1, 4, 4}},
			1,
		},
		{
			`Test 3`,
			args{target: 11, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			0,
		},
		{
			`Test 4`,
			args{target: 1, nums: []int{}},
			0,
		}, {
			`Test 5`,
			args{target: 3, nums: []int{1, 1, 1, 1, 1, 1, 1, 1}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
