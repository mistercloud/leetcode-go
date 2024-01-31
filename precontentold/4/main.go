package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(read, writer)
}

type Void struct {
}

type Friend struct {
	numCards int
	pos      int
}

func Solution(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	var totalFriends, totalCards int
	fmt.Fscanln(reader, &totalFriends, &totalCards)
	friends := make([]Friend, totalFriends)
	for i := 0; i < totalFriends; i++ {
		var numCards int
		fmt.Fscan(reader, &numCards)
		friends[i] = Friend{numCards, i}
	}

	cards := make(map[int]Void, totalCards)
	for i := 1; i <= totalCards; i++ {
		cards[i] = Void{}
	}

	//сортируем друзей по количеству карт
	sort.Slice(friends, func(i, j int) bool {
		return friends[i].numCards >= friends[j].numCards
	})

	gives := make([]int, totalFriends)
	for _, friend := range friends {
		found := false
		for i := totalCards; i > friend.numCards; i-- {
			_, ok := cards[i]
			if ok {
				found = true
				gives[friend.pos] = i
				delete(cards, i)
				totalCards--
				break
			}
		}
		if !found {
			gives[0] = -1
			break
		}
	}
	fmt.Println(friends)
	if gives[0] != -1 {
		for _, give := range gives {
			fmt.Fprint(writer, give, " ")
		}
	} else {
		fmt.Fprint(writer, -1)
	}
	fmt.Fprintln(writer)
}
