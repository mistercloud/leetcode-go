package main

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] != ' ' {
			for i >= 0 && s[i] != ' ' {
				cnt++
				i--
			}
			return cnt
		}
	}

	return -1
}
