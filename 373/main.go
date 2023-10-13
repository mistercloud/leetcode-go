package main

import "container/heap"

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {

	h := &SumHeap{}
	heap.Init(h)
	for i := range nums1 {
		heap.Push(h, SumItem{nums1[i] + nums2[0], i, 0})
	}

	ret := make([][]int, 0)
	for !h.Empty() && k > 0 {

		item := heap.Pop(h).(SumItem)
		ret = append(ret, []int{nums1[item.i], nums2[item.j]})
		nextElement := item.j + 1
		if nextElement < len(nums2) {
			heap.Push(h, SumItem{sum: nums1[item.i] + nums2[nextElement], i: item.i, j: nextElement})
		}

		k--
	}
	return ret
}

type SumItem struct {
	sum int //сумма элементов
	i   int //позиция в первом массиве
	j   int //позиция в втором массиве
}

type SumHeap []SumItem

func (h SumHeap) Len() int           { return len(h) }
func (h SumHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h SumHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h SumHeap) Empty() bool        { return len(h) == 0 }

func (h *SumHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(SumItem))
}

func (h *SumHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
