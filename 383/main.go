package main

import (
	"fmt"
)

func main() {
	ransomNote := `a`
	magazine := `b`
	fmt.Println(ransomNote, ` `, magazine, ` - `, canConstruct(ransomNote, magazine))

	ransomNote = `aa`
	magazine = `ab`
	fmt.Println(ransomNote, ` `, magazine, ` - `, canConstruct(ransomNote, magazine))

	ransomNote = `aa`
	magazine = `aab`
	fmt.Println(ransomNote, ` `, magazine, ` - `, canConstruct(ransomNote, magazine))

	ransomNote = `aab`
	magazine = `a`
	fmt.Println(ransomNote, ` `, magazine, ` - `, canConstruct(ransomNote, magazine))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(magazine) < len(ransomNote) {
		return false
	}
	letters := make(map[rune]int)
	for _, char := range magazine {
		letters[char]++
	}
	for _, char := range ransomNote {
		letters[char]--
		if letters[char] < 0 {
			return false
		}
	}
	return true

	/*for _, v := range ransomNote {
		if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v)) {
			return false
	}*/
}
