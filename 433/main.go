package main

import (
	"fmt"
	"math"
)

func main() {
	startGene := `AACCGGTT`
	endGene := `AACCGGTA`
	bank := []string{`AACCGGTA`}
	fmt.Println(startGene, endGene, minMutation(startGene, endGene, bank))

	startGene = `AACCGGTT`
	endGene = `AAACGGTA`
	bank = []string{`AACCGGTA`, `AACCGCTA`, `AAACGGTA`}
	fmt.Println(startGene, endGene, minMutation(startGene, endGene, bank))
}

func minMutation(startGene string, endGene string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}
	//не нужно мутаций
	if startGene == endGene {
		return 0
	}

	dp := make(map[string]int, len(bank)+1)
	for _, gene := range bank {
		dp[gene] = math.MaxInt
	}

	_, ok := dp[endGene]
	//в банке не содержится финальный ген, невозможно прийти
	if !ok {
		return -1
	}

	//BFS queue
	q := make([]string, 0)
	q = append(q, startGene)
	dp[startGene] = 0
	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		for _, gen := range bank {
			if CanMutate(current, gen) {
				if dp[gen] > dp[current]+1 {
					q = append(q, gen)
					dp[gen] = dp[current] + 1
				}
			}
		}
	}

	fmt.Println(dp)
	if dp[endGene] == math.MaxInt {
		return -1
	}
	return dp[endGene]
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
