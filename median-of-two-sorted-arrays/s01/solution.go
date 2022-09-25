package s01

import (
	"sort"
)

// From Accepted Solutions Runtime Distribution - sample 3 ms submission
// https://leetcode.com/submissions/api/detail/4/golang/3/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    	var result float64

	newSortArray := append(nums1, nums2...)
	sort.Ints(newSortArray)

	nsLen := len(newSortArray)
	median := nsLen / 2
	remainder := nsLen % 2

	if remainder == 0 {
		result = float64(newSortArray[median-1]+newSortArray[median]) / 2.0
		return result
	} else {
		return float64(newSortArray[median])
	}
}


