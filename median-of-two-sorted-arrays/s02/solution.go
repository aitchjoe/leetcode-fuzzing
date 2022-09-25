package s02

import (
	"math"
)

// From Accepted Solutions Memory Distribution - sample 5000 KB submission
// https://leetcode.com/submissions/api/detail/4/golang/memory/5000/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    l := (m + n + 1) / 2 
    r := (m + n + 2) / 2 
    val1 := float64(kthLargest(nums1, nums2, l)) 
    val2 := float64(kthLargest(nums1, nums2, r))
    //fmt.Println(val1, val2)
    
    return (val1+val2)/float64(2)
}

func kthLargest(a, b []int, k int) int {
    if len(a) == 0 {
        return b[k-1]
    }
    if len(b) == 0 {
        return a[k-1]
    }
    
    if k == 1 {
        return min(a[0], b[0]) 
    }
    
    mid := k/2-1
    
    aVal := math.MaxInt
    if mid < len(a) {
       aVal = a[mid] 
    }
    
    bVal := math.MaxInt
    if mid < len(b) {
        bVal = b[mid]
    }
    
    if aVal < bVal {
        return kthLargest(a[mid+1:], b, k-k/2)
    }
    return kthLargest(a, b[mid+1:], k-k/2) 
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
