package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	start, end := 0, 0
	maxLen := 1
	for end < len(s) {
		//не имеет смысла проверять текущее окно - расширяем вправо
		if end-start+1 > maxLen {
			if IsUniq(s[start : end+1]) {
				maxLen = end - start + 1
			} else {
				//двигаем все окно вправо
				start++
			}
		}

		end++

	}
	return maxLen
}

func IsUniq(s string) bool {
	hmap := make(map[rune]bool, len(s))
	for _, ch := range s {
		if _, ok := hmap[ch]; ok == false {
			hmap[ch] = true
		} else {
			return false
		}
	}

	return true
}
