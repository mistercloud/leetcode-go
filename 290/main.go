package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/word-pattern/
func main() {
	s := `dog cat cat dog`
	p := `abba`
	fmt.Println(s, p, wordPattern(p, s))

	s = `dog cat cat fish`
	p = `abba`
	fmt.Println(s, p, wordPattern(p, s))

	s = `dog cat cat dog`
	p = `aaaa`
	fmt.Println(s, p, wordPattern(p, s))

	s = `dog cat bbbb aaaa`
	p = `acde`
	fmt.Println(s, p, wordPattern(p, s))
}

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, ` `)
	if len(pattern) != len(words) {
		return false
	}
	hmap := map[string]byte{}
	hmapChars := map[byte]bool{}
	for i := 0; i < len(words); i++ {
		if hmap[words[i]] == 0 && !hmapChars[pattern[i]] {
			hmap[words[i]] = pattern[i]
			hmapChars[pattern[i]] = true
		}
		if hmap[words[i]] != pattern[i] {
			return false
		}
	}
	return true
}
