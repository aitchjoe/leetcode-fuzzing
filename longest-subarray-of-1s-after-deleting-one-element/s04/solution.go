package s04

// From Accepted Solutions Runtime Distribution - sample 99 ms submission
// https://leetcode.com/submissions/api/detail/1586/golang/99/

func LongestSubarray(nums []int) int {
	maxLen := 0
	last, current := 0, 0

	for _, v := range nums {
		if v == 0 {
			maxLen = max(maxLen, last+current)
			last = current
			current = 0
		} else {
			current++
		}
	}
	maxLen = max(maxLen, last+current)

	if maxLen == len(nums) {
		return maxLen - 1
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
