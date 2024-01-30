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

	var numTests int
	fmt.Fscanln(reader, &numTests)
	for numTests > 0 {
		var dataNum int
		fmt.Fscanln(reader, &dataNum)
		minMax := []int{15, 30}
		for dataNum > 0 {
			var operation string
			var temp int
			fmt.Fscanln(reader, &operation, &temp)
			if minMax[0] > minMax[1] {
				fmt.Fprintln(writer, "-1")
				dataNum--
				continue
			}
			if operation == ">=" {
				if temp > minMax[0] {
					minMax[0] = temp
				}
			} else {
				if temp < minMax[1] {
					minMax[1] = temp
				}

			}
			if minMax[0] <= minMax[1] {
				fmt.Fprintln(writer, minMax[0])
			} else {
				fmt.Fprintln(writer, "-1")
			}
			dataNum--
		}
		fmt.Fprint(writer, "\n")
		numTests--
	}

}
