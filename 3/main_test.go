package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test1",
			args{"abcabcbb"},
			3,
		},
		{
			"Test2",
			args{"bbbbb"},
			1,
		},
		{
			"Test3",
			args{"pwwkew"},
			3,
		},
		{
			"Test4",
			args{""},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUniq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test1",
			args{"a"},
			true,
		},
		{
			"Test2",
			args{"ab"},
			true,
		},
		{
			"Test2",
			args{"aba"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniq(tt.args.s); got != tt.want {
				t.Errorf("IsUniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
