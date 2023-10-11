package main

func longestCommonPrefix(strs []string) string {
	i := 0
	for i = 0; i < len(strs[0]); i++ {
		for wordIndex, word := range strs {
			if i >= len(word) {
				return strs[0][0:i]
			}
			if wordIndex != 0 && word[i] != strs[0][i] {
				return strs[0][0:i]
			}
		}

	}
	return strs[0][0:i]
}
