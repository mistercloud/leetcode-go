package main

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"Test1",
			args{43261596},
			964176192,
		},
		{
			"Test2",
			args{4294967293},
			3221225471,
		},
		{
			"Test3",
			args{1},
			2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
