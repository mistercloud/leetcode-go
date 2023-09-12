package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			`Test 1`,
			args{s1: `this apple is sweet`, s2: `this apple is sour`},
			[]string{`sweet`, `sour`},
		},
		{
			`Test 2`,
			args{s1: `apple apple`, s2: `banana`},
			[]string{`banana`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func uncommonFromSentences(s1 string, s2 string) []string {
	s := s1 + ` ` + s2
	words := strings.Split(s, ` `)
	wordsCnt := make(map[string]int)
	for _, word := range words {
		wordsCnt[word]++
	}
	var uncommon []string

	for word, cnt := range wordsCnt {
		if cnt == 1 {
			uncommon = append(uncommon, word)
		}
	}
	return uncommon
}
