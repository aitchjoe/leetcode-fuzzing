package s05

// From Accepted Solutions Memory Distribution - sample 8700 KB submission
// https://leetcode.com/submissions/api/detail/1586/golang/memory/8700/

func LongestSubarray(nums []int) int {
	curslicelen, prevslicelen, max := 0, 0, 0
	prev, haszeros := false, false
	for _, num := range nums {
		if num == 1 {
			curslicelen++
			if max < curslicelen+prevslicelen {
				max = curslicelen + prevslicelen
			}
			prev = false
		}
		if num == 0 {
			if !haszeros {
				haszeros = true
			}
			if prev {
				prevslicelen = 0
			} else {
				prevslicelen = curslicelen
			}
			curslicelen = 0
			prev = true
		}
	}

	if !haszeros {
		max--
	}

	return max
}
