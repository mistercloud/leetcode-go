package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/
func main() {
	x := []int{1, 3}
	y := []int{2}
	fmt.Println(x, y, findMedianSortedArrays(x, y), naiveMedian(x, y))

	x = []int{1, 2}
	y = []int{3, 4}
	fmt.Println(x, y, findMedianSortedArrays(x, y), naiveMedian(x, y))
}

// The idea is simple:
// 0. we know the median should be at index (m+n)/2
// 1. odd or even, finding this index solves the problem
// 2. we use binary search for the smaller array
// 3. first iteration we find the half way of nums1 of index h1
// 4. to make the half of the overall, we take (m+n)/2-h1 from nums2 to be h2
// 5. define l1, r1, l2, r2 to be the left & right of nums1 & nums2 around the h1 & h2
// 6. if l1<r2 && l2<r1, then we found the correct indexes in both array
// 7. return the result based of odd or even of the result size
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	start := 0
	end := len(nums1)
	midInMerged := (len(nums1) + len(nums2) + 1) / 2

	for start <= end {
		mid := (start + end) / 2
		firstLeftSize := mid
		secondLeftSize := midInMerged - mid
		leftFirst := math.MinInt
		if firstLeftSize > 0 {
			leftFirst = nums1[firstLeftSize-1]
		}
		leftSecond := math.MinInt
		if secondLeftSize > 0 {
			leftSecond = nums2[secondLeftSize-1]
		}
		rightFirst := math.MaxInt
		if firstLeftSize < len(nums1) {
			rightFirst = nums1[firstLeftSize]
		}
		rightSecond := math.MaxInt
		if secondLeftSize < len(nums2) {
			rightSecond = nums2[secondLeftSize]
		}
		if leftFirst <= rightSecond && leftSecond <= rightFirst {
			//четный мерж
			if (len(nums1)+len(nums2))%2 == 0 {
				return (math.Max(float64(leftFirst), float64(leftSecond)) +
					math.Min(float64(rightFirst), float64(rightSecond))) / 2
			}
			return math.Max(float64(leftFirst), float64(leftSecond))
		} else if leftFirst > rightSecond {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0
}

// наивный метод, для проверки значения
func naiveMedian(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	if len(nums1)%2 == 1 {
		return float64(nums1[len(nums1)/2])
	}

	mid := len(nums1) / 2
	return (float64(nums1[mid-1]) + float64(nums1[mid])) / 2
}
