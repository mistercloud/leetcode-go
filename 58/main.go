package main

import "fmt"

func main() {
	s := `Hello World`
	fmt.Println(s, ` `, lengthOfLastWord(s))

	s = `   fly me   to   the moon  `
	fmt.Println(s, ` `, lengthOfLastWord(s))

	s = `luffy is still joyboy`
	fmt.Println(s, ` `, lengthOfLastWord(s))

	s = ``
	fmt.Println(s, ` `, lengthOfLastWord(s))
}

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	cnt := 0
	meetChar := false
	for i := len(s) - 1; i >= 0; i-- {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			meetChar = true
			cnt++
		}
		if s[i] == ' ' && meetChar {
			break
		}
	}

	return cnt
}
