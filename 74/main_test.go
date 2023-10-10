package main

import "testing"

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test1",
			args{[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			}, 3},
			true,
		},
		{
			"Test2",
			args{[][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			}, 13},
			false,
		},
		{
			"Test3",
			args{[][]int{
				{1},
			}, 13},
			false,
		},
		{
			"Test4",
			args{[][]int{
				{1},
			}, 1},
			true,
		},
		{
			"Test5",
			args{[][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
			}, 10},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
