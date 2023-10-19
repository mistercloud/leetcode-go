package _5

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test1",
			args{"aaaaa"},
			"aaaaa",
		},
		{
			"Test1",
			args{"babad"},
			"bab",
		},
		{
			"Test2",
			args{"cbbd"},
			"bb",
		},
		{
			"Test3",
			args{"a"},
			"a",
		},
		{
			"Test4",
			args{"aaaaa"},
			"aaaaa",
		},
		{
			"Test5",
			args{"aaaaab"},
			"aaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
