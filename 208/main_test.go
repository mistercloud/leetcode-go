package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Trie
	}{
		{
			`Test 1`,
			Trie{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_Insert(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			`Test 1`,
			args{`apple`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor()
			this.Insert(tt.args.word)
		})
	}
}

func TestTrie_Search(t *testing.T) {

	node := Constructor()
	node.Insert(`apple`)
	node.Insert(`banana`)
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{"apple"},
			true,
		},
		{
			"Test 2",
			args{"banana"},
			true,
		},
		{
			"Test 3",
			args{"ap"},
			false,
		},
		{
			"Test 4",
			args{"pear"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := node.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrie_StartsWith(t *testing.T) {
	node := Constructor()
	node.Insert(`apple`)
	node.Insert(`banana`)
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{"apple"},
			true,
		},
		{
			"Test 2",
			args{"banana"},
			true,
		},
		{
			"Test 3",
			args{"app"},
			true,
		},
		{
			"Test 4",
			args{"pear"},
			false,
		},
		{
			"Test 4",
			args{"banar"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := node.StartsWith(tt.args.prefix); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
