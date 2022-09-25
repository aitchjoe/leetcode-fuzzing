package s01

// From Accepted Solutions Runtime Distribution - sample 30 ms submission
// https://leetcode.com/submissions/api/detail/1586/golang/30/

func LongestSubarray(nums []int) int {
	left, right, zeros, ans := 0, 0, 0, 0

	anyZeros := false
	maxScore := 0
	for left <= right && right < len(nums) {
		if nums[right] == 0 && zeros < 2 {
			zeros++
			anyZeros = true
			right++
		} else if nums[right] == 1 {
			right++
			ans++
			maxScore = max(maxScore, ans)
		}
		if zeros == 2 {
			if nums[left] == 0 {
				left++
				anyZeros = true
				zeros--
			} else {
				ans--
				left++
			}
		}
	}

	if !anyZeros {
		return maxScore - 1
	}

	return maxScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
