package _198

import "testing"

func Test_rob(t *testing.T) {
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
			args{[]int{1, 2, 3, 1}},
			4,
		},
		{
			"Test2",
			args{[]int{2, 7, 9, 3, 1}},
			12,
		},
		{
			"Test3",
			args{[]int{1, 1, 9, 1, 2, 1}},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
