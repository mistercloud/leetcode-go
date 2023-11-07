package _242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	smap := make(map[rune]int, len(s))
	for _, ch := range s {
		smap[ch]++
	}

	for _, ch := range t {
		smap[ch]--
		if smap[ch] < 0 {
			return false
		}
	}

	return true
}
