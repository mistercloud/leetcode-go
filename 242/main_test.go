package _242

import "testing"

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test1",
			args{"anagram", "nagaram"},
			true,
		},
		{
			"Test2",
			args{"rat", "car"},
			false,
		},
		{
			"Test3",
			args{"anagram", "nagaramwe"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
