package main

// https://leetcode.com/problems/text-justification/
func fullJustify(words []string, maxWidth int) []string {
	var ret []string

	for len(words) > 0 {
		curLen := 0
		pos := 0
		curLen += len(words[pos])
		for {
			if pos+1 == len(words) {
				//не осталось слов
				break
			}
			//добавляем слово
			curLen += len(words[pos+1]) + 1
			if curLen > maxWidth {
				break

			}
			pos++
		}
		pos++
		sub := words[0:pos]
		words = words[pos:]
		if len(words) > 0 {
			ret = append(ret, justifyCenter(sub, maxWidth))
		} else {
			ret = append(ret, justifyLeft(sub, maxWidth))
		}
	}
	return ret
}

func justifyLeft(words []string, maxWidth int) string {
	ret := make([]rune, maxWidth)
	i := 0
	for _, word := range words {
		for _, ch := range word {
			ret[i] = ch
			i++
		}
		if i < maxWidth {
			ret[i] = ' '
		}

		i++
	}
	for i < maxWidth {
		ret[i] = ' '
		i++
	}
	return string(ret)
}

func justifyCenter(words []string, maxWidth int) string {
	ret := make([]rune, maxWidth)
	for i := range ret {
		ret[i] = ' '
	}
	//ставим первое слово и последние если есть
	start := 0
	end := len(ret) - 1
	for _, ch := range words[0] {
		ret[start] = ch
		start++
	}
	if len(words) > 1 {
		words = words[1:]
		if len(words) > 0 {
			for i, ch := range words[len(words)-1] {
				ret[end-len(words[len(words)-1])+i*2+1] = ch
				end--
			}
		}
		words = words[:len(words)-1]
		if len(words) > 0 {
			//еще остались слова
			wordsLen := 0
			for _, word := range words {
				wordsLen += len(word)
			}
			intervals := make([]int, len(words)+1)
			j := 0
			for i := 0; i <= end-start-wordsLen; i++ {
				intervals[j]++
				j++
				if j > len(intervals)-1 {
					j = 0
				}
			}
			for _, interval := range intervals {
				for i := 0; i < interval; i++ {
					ret[start] = ' '
					start++
				}
				if len(words) > 0 {
					for _, ch := range words[0] {
						ret[start] = ch
						start++
					}
					words = words[1:]
				}

			}
		}
	}

	return string(ret)
}
