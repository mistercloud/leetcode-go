package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 3
	fmt.Println(x, ` - `, fizzBuzz(x))

	x = 5
	fmt.Println(x, ` - `, fizzBuzz(x))

	x = 15
	fmt.Println(x, ` - `, fizzBuzz(x))
}

func fizzBuzz(n int) []string {

	result := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			result[i-1] += `Fizz`
		}
		if i%5 == 0 {
			result[i-1] += `Buzz`
		}
		if len(result[i-1]) == 0 {
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}
