package solution

import (
	"testing"
)

type TestCase struct {
	Input []int
	Want  int
}

func TestLongestSubarray(t *testing.T) {
	testcases := []TestCase{
		{
			Input: []int{1, 1, 0, 1},
			Want:  3,
		},
		{
			Input: []int{1, 1, 0, 1},
			Want:  3,
		},
		{
			Input: []int{1, 1, 0, 1},
			Want:  3,
		},

		// my accepted submission failed by this case, leetcode have not enough test cases for this problem.
		{
			Input: []int{1, 1, 1, 0},
			Want:  3,
		},
	}
	for _, tc := range testcases {
		out := longestSubarray(tc.Input)
		if out != tc.Want {
			t.Fatalf(`longestSubarray(%v) = %d, want %d`, tc.Input, out, tc.Want)
		}
	}
}
