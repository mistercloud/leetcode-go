package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := `11`
	s2 := `1`
	expected := `100`
	fmt.Println(s1, s2, addBinary(s1, s2), expected)

	s1 = `1010`
	s2 = `1011`
	expected = `10101`
	fmt.Println(s1, s2, addBinary(s1, s2), expected)

}

// https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1

	carry := 0
	sum := 0
	ret := ``
	for i >= 0 && j >= 0 {
		first := int(a[i] - '0')
		second := int(b[j] - '0')

		sum, carry = sumBinary(first, second, carry)

		ret = strconv.Itoa(sum) + ret
		i--
		j--

	}
	for i >= 0 {
		first := int(a[i] - '0')
		sum, carry = sumBinary(first, 0, carry)
		ret = strconv.Itoa(sum) + ret
		i = i - 1
	}

	for j >= 0 {
		second := int(b[j] - '0')

		sum, carry = sumBinary(0, second, carry)

		ret = strconv.Itoa(sum) + ret
		j = j - 1
	}

	if carry > 0 {
		ret = strconv.Itoa(1) + ret
	}

	return ret
}

func sumBinary(a, b, carry int) (int, int) {
	ret := a + b + carry
	switch ret {
	case 3:
		return 1, 1
	case 2:
		return 0, 1
	case 1:
		return 1, 0
	}

	return 0, 0
}
