package main

import "fmt"

func main() {
	s := `A man, a plan, a canal: Panama`
	fmt.Println(s, ` - `, isPalindrome(s))

	s = `1221`
	fmt.Println(s, ` - `, isPalindrome(s))

	s = `race a car`
	fmt.Println(s, ` - `, isPalindrome(s))

	s = ` `
	fmt.Println(s, ` - `, isPalindrome(s))
}

// https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {

	arr := []rune{}
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char += ' '
		}
		if (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z') {
			arr = append(arr, char)
		}
	}
	j := len(arr) - 1
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[j] {
			return false
		}
		j--
	}

	return true
}
