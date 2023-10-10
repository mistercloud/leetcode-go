package main

import "testing"

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{[]int{2, 2, 1}},
			1,
		},
		{
			"Test2",
			args{[]int{4, 1, 2, 1, 2}},
			4,
		},
		{
			"Test3",
			args{[]int{1}},
			1,
		},
		{
			"Test4",
			args{[]int{1, 3, 5, 7, 5, 3, 1}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
