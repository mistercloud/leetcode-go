package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test1",
			args{[]string{"flower", "flow", "flight"}},
			"fl",
		},
		{
			"Test2",
			args{[]string{"dog", "racecar", "car"}},
			"",
		},
		{
			"Test3",
			args{[]string{"dog", "", "car"}},
			"",
		},
		{
			"Test4",
			args{[]string{"", "flow", "flight"}},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
