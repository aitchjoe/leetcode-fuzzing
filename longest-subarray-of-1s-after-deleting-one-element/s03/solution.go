package s03

// From Accepted Solutions Runtime Distribution - sample 34 ms submission
// https://leetcode.com/submissions/api/detail/1586/golang/34/

func LongestSubarray(nums []int) int {
	res := 0
	curr := 0
	priv := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			curr++
			continue
		}

		if i+1 < len(nums) && nums[i+1] != 0 {
			r := curr + priv
			res = max(res, r)
			priv = curr
			curr = 0

			continue
		}

		r := curr + priv
		res = max(res, r)
		curr = 0
		priv = 0
	}

	r := curr + priv
	res = max(res, r)

	if res == len(nums) {
		res--
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
