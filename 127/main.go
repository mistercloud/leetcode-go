package main

import (
	"fmt"
	"math"
)

func main() {
	startGene := `hit`
	endGene := `cog`
	bank := []string{`hot`, `dot`, `dog`, `lot`, `log`, `cog`}
	fmt.Println(startGene, endGene, ladderLength(startGene, endGene, bank))

	startGene = `hit`
	endGene = `cog`
	bank = []string{`hot`, `dot`, `dog`, `lot`, `log`}
	fmt.Println(startGene, endGene, ladderLength(startGene, endGene, bank))
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}

	dp := make(map[string]int, len(wordList)+1)
	for _, word := range wordList {
		dp[word] = math.MaxInt
	}

	_, ok := dp[endWord]
	//в словаре нет финального слова
	if !ok {
		return 0
	}

	//BFS queue
	q := make([]string, 0)
	q = append(q, beginWord)
	dp[beginWord] = 1
	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		for _, word := range wordList {
			if CanMutate(current, word) {
				if dp[word] > dp[current]+1 {
					q = append(q, word)
					dp[word] = dp[current] + 1
				}
			}
		}
	}

	fmt.Println(dp)
	if dp[endWord] == math.MaxInt {
		return 0
	}
	return dp[endWord]
}

func CanMutate(start, end string) bool {
	if start == end {
		return false
	}
	mutations := 0
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			mutations++
			if mutations > 1 {
				return false
			}
		}

	}

	return mutations == 1
}
