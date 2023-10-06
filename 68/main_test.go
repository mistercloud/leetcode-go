package main

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			`Test1`,
			args{
				[]string{`This`, `is`, `an`, `example`, `of`, `text`, `justification.`},
				16,
			},
			[]string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			`Test2`,
			args{
				[]string{"What", "must", "be", "acknowledgment", "shall", "be"},
				16,
			},
			[]string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			`Test3`,
			args{
				[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				20,
			},
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_justifyLeft(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			``,
			args{[]string{"a", "word."}, 10},
			`a word.   `,
		},
		{
			``,
			args{[]string{"a", "b"}, 10},
			`a b       `,
		},
		{
			``,
			args{[]string{"abc", "dfg", "12"}, 10},
			`abc dfg 12`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := justifyLeft(tt.args.words, tt.args.maxWidth); got != tt.want {
				t.Errorf("justifyLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_justifyCenter(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			``,
			args{[]string{"a", "word."}, 10},
			`a    word.`,
		},
		{
			``,
			args{[]string{"a", "b"}, 10},
			`a        b`,
		},
		{
			``,
			args{[]string{"abc", "dfg", "12"}, 10},
			`abc dfg 12`,
		},
		{
			``,
			args{[]string{"a", "dfg", "b"}, 10},
			`a   dfg  b`,
		},
		{
			``,
			args{[]string{`This`, `is`, `an`}, 16},
			`This    is    an`,
		},
		{
			``,
			args{[]string{`acknowledgment`}, 16},
			`acknowledgment  `,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := justifyCenter(tt.args.words, tt.args.maxWidth); got != tt.want {
				t.Errorf("justifyCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}
