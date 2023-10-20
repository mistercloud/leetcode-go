package _45

import "testing"

func Test_jump(t *testing.T) {
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
			args{[]int{2, 3, 1, 1, 4}},
			2,
		},
		{
			"Test2",
			args{[]int{2, 3, 0, 1, 4}},
			2,
		},
		{
			"Test3",
			args{[]int{4, 3, 0, 1, 4}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
