package main

import "fmt"

func findMedian(nums1 []int, nums2 []int, step int) int {
	fmt.Println("nums1 ", nums1, " nums2 ", nums2)
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findMedian(nums2, nums1, step)
	}
	if len1 == 0 {
		return nums2[step-1]
	}
	if step == 1 {
		if nums2[step-1] > nums1[step-1] {
			return nums1[step-1]
		} else {
			return nums2[step-1]
		}
	}
	var part1 int
	if step/2 > len1 {
		part1 = len1
	} else {
		part1 = step / 2
	}
	part2 := step - part1
	fmt.Println("part1 ", part1, " part2 ", part2)
	if nums1[part1-1] == nums2[part2-1] {
		return nums1[part1-1]
	} else if nums1[part1-1] > nums2[part2-1] {
		return findMedian(nums1, nums2[part2:], step-part2)
	} else {
		return findMedian(nums1[part1:], nums2, step-part1)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	sumlen := len1 + len2
	fmt.Println(len1)
	fmt.Println(len2)
	fmt.Println(sumlen)

	fmt.Println("")
	var step int = sumlen / 2
	fmt.Println(step)
	if sumlen%2 == 1 {
		return float64(findMedian(nums1, nums2, step+1))
	} else {
		return (float64(findMedian(nums1, nums2, step)) +
			float64(findMedian(nums1, nums2, step+1))) / 2.0
	}
}

func main() {
	a := []int{1, 3}
	b := []int{2}
	var ret float64 = findMedianSortedArrays(a, b)
	fmt.Println(ret)
}
