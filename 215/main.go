package main

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {

	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	for {
		k--
		if k == 0 {
			return heap.Pop(h).(int)
		}
		heap.Pop(h)
	}
}

//решение через реализацию heap
/**func findKthLargest(nums []int, k int) int {

	for i := len(nums) / 2; i >= 0; i-- {
		heapify(nums, i)
	}

	for {
		k--
		if k == 0 {
			return nums[0]
		}
		nums[0] = nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		heapify(nums, 0)
	}

	return -1
}

func heapify(nums []int, i int) {

	for {
		leftIndex := 2*i + 1
		rightIndex := 2*i + 2
		largest := i
		if leftIndex < len(nums) && nums[leftIndex] > nums[largest] {
			largest = leftIndex
		}
		if rightIndex < len(nums) && nums[rightIndex] > nums[largest] {
			largest = rightIndex
		}
		if largest == i {
			break
		}

		nums[i], nums[largest] = nums[largest], nums[i]
		i = largest
	}

}**/
