package main

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
			args{"Hello World"},
			5,
		},
		{
			"Test2",
			args{"   fly me   to   the moon  "},
			4,
		},
		{
			"Test3",
			args{"luffy is still joyboy"},
			6,
		},
		{
			"Test4",
			args{"  "},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
