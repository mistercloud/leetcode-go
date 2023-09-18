package main

import "fmt"

func main() {
	n1, n2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6, 5, 2, 5}
	n, m := 3, 3
	fmt.Println(n1, ` `, n2, ` `, n, ` `, m)
	merge(n1, n, n2, m)
	fmt.Println(n1)

	n1, n2 = []int{1}, []int{}
	n, m = 1, 0
	fmt.Println(n1, ` `, n2, ` `, n, ` `, m)
	merge(n1, n, n2, m)
	fmt.Println(n1)

	n1, n2 = []int{0}, []int{1}
	n, m = 0, 1
	fmt.Println(n1, ` `, n2, ` `, n, ` `, m)
	merge(n1, n, n2, m)
	fmt.Println(n1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := i + n
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
