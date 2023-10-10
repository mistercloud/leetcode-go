package main

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {

	digiMap := [][]string{
		{`a`, `b`, `c`},
		{`d`, `e`, `f`},
		{`g`, `h`, `i`},
		{`j`, `k`, `l`},
		{`m`, `n`, `o`},
		{`p`, `q`, `r`, `s`},
		{`t`, `u`, `v`},
		{`w`, `x`, `y`, `z`},
	}

	ret := []string{}
	for _, digi := range digits {
		if len(ret) == 0 {
			ret = digiMap[digi-'2']
		} else {
			ret = merge(&ret, &digiMap[digi-'2'])
		}

	}

	return ret
}

func merge(ret, chars *[]string) []string {

	var result []string
	for _, retItem := range *ret {
		for _, ch := range *chars {
			result = append(result, retItem+ch)
		}
	}

	return result

}
