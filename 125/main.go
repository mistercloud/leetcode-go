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
		if char >= 65 && char <= 90 {
			char += 32
		}
		if (char >= 48 && char <= 57) || (char >= 97 && char <= 122) {
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
