package s02

// From Accepted Solutions Memory Distribution - sample 7200 KB submission
// https://leetcode.com/submissions/api/detail/1586/golang/memory/7200/

func LongestSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	beforeZeroes := make([]int, 0)

	counter := 0

	for _, num := range nums {
		if num == 1 {
			counter++
		} else {
			beforeZeroes = append(beforeZeroes, counter)
			counter = 0
		}
	}

	if counter > 0 {
		beforeZeroes = append(beforeZeroes, counter)
	}

	if len(beforeZeroes) == 1 {
		return beforeZeroes[0] - 1
	}

	result := 0

	for i := 1; i < len(beforeZeroes); i++ {
		sum := beforeZeroes[i-1] + beforeZeroes[i]
		if sum > result {
			result = sum
		}
	}

	return result
}
