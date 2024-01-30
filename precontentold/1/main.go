package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	Solution(read, writer)
}

func Solution(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	var ret []byte
	fmt.Fscanln(reader, &ret)
	var numRewrites int
	fmt.Fscanln(reader, &numRewrites)
	for i := 0; i < numRewrites; i++ {
		var start, end int
		var rewrite string
		fmt.Fscanln(reader, &start, &end, &rewrite)
		k := 0
		for k+start <= end {
			ret[k+start-1] = rewrite[k]
			k++
		}
	}
	fmt.Fprintln(writer, string(ret))
}
