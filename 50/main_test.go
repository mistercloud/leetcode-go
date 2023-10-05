package main

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			`Test 1`,
			args{2.00000, 10},
			1024.00000,
		},
		{
			`Test 2`,
			args{2.10000, 3},
			9.26100,
		},
		{
			`Test 3`,
			args{2.00000, -2},
			0.25000,
		},
		{
			`Test 4`,
			args{0, -2},
			0,
		},
		{
			`Test 5`,
			args{1, 0},
			1,
		},
		{
			`Test 6`,
			args{2, 2},
			4,
		},
		{
			`Test 7`,
			args{0.44528, 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
