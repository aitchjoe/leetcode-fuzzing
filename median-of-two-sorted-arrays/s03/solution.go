package s03

// From Accepted Solutions Memory Distribution - sample 6600 KB submission
// https://leetcode.com/submissions/api/detail/4/golang/memory/6600/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    ints := merge(nums1, nums2)
	length := len(ints)
	if length%2 == 0 {
		return float64(ints[length/2-1]+ints[length/2]) / 2
	}
	return float64(ints[length/2])
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
