package main

import (
	"sort"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {

	hmap := make(map[string][]string, len(strs))
	for _, word := range strs {
		sb := []byte(word)
		sort.Slice(sb, func(i, j int) bool {
			return sb[i] < sb[j]
		})
		hmap[string(sb)] = append(hmap[string(sb)], word)

	}
	var ret [][]string
	for _, anagrams := range hmap {
		ret = append(ret, anagrams)
	}
	return ret
}
