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

var (
	BLANK byte = '.'
)

func Solution(reader *bufio.Reader, writer *bufio.Writer) {
	defer writer.Flush()

	var numTests int
	fmt.Fscanln(reader, &numTests)
	for numTests > 0 {
		var frameNum, frameHeight, frameWidth int
		fmt.Fscanln(reader, &frameNum, &frameHeight, &frameWidth)
		frames := make([][][]byte, frameNum)

		for i := 0; i < frameNum; i++ {
			frame := make([][]byte, frameHeight)
			for j := 0; j < frameHeight; j++ {
				row := make([]byte, frameWidth)
				fmt.Fscanln(reader, &row)
				frame[j] = row
			}

			frames[i] = frame
			if i < frameNum-1 {
				fmt.Fscanln(reader)
			}

		}
		concatFrames(frames)
		for i := 0; i < len(frames[0]); i++ {
			for j := 0; j < len(frames[0][0]); j++ {
				fmt.Fprint(writer, string(frames[0][i][j]))
			}
			fmt.Fprintln(writer)
		}
		fmt.Fprint(writer, "\n")

		numTests--
	}

}

func concatFrames(frames [][][]byte) {

	for frame := 1; frame < len(frames); frame++ {
		for i := 0; i < len(frames[0]); i++ {
			for j := 0; j < len(frames[0][0]); j++ {
				if frames[frame][i][j] != BLANK && frames[0][i][j] == BLANK {
					frames[0][i][j] = frames[frame][i][j]
				}
			}
		}
	}
}
