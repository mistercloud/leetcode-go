package main

import "fmt"

func main() {
	h := `sadbutsad`
	n := `sad`
	fmt.Println(h, n, strStr(h, n))

	h = `leetcode`
	n = `leeto`
	fmt.Println(h, n, strStr(h, n))

	h = `aaaaasadbutsad`
	n = `sad`
	fmt.Println(h, n, strStr(h, n))

	h = `l`
	n = `leeto`
	fmt.Println(h, n, strStr(h, n))

	h = `leeto`
	n = `leeto`
	fmt.Println(h, n, strStr(h, n))

	h = `abc`
	n = `c`
	fmt.Println(h, n, strStr(h, n))
}

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if needle == haystack {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
