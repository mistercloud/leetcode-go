package main

import "fmt"

func main() {
	num := 14
	fmt.Println(num, ` - `, numberOfSteps(num))
	num = 8
	fmt.Println(num, ` - `, numberOfSteps(num))
	num = 123
	fmt.Println(num, ` - `, numberOfSteps(num))
	num = 0
	fmt.Println(num, ` - `, numberOfSteps(num))

}

func numberOfSteps(num int) int {
	i := 0
	for ; num > 0; i++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
	}
	return i
}
