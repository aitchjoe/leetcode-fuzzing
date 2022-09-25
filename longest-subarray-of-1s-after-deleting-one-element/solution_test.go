package solution

import (
	"math/rand"
	"testing"
	"time"

	// [proposal: Go 2: cmd/go: allow relative imports](https://github.com/golang/go/issues/20883)
	"github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element/s01"
	"github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element/s02"
	"github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element/s03"
	"github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element/s04"
	"github.com/aitchjoe/leetcode-fuzzing/longest-subarray-of-1s-after-deleting-one-element/s05"
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

func FuzzLongestSubarray(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(10)
	f.Add(25600)
	f.Fuzz(func(t *testing.T, numsLength int) {
		// 1 <= nums.length <= 10^5
		n := numsLength
		if n < 0 {
			n = -n
		}
		n = n%100000 + 1

		arr := make([]int, n, n)
		for i := 0; i < n; i++ {
			arr[i] = int(rand.Uint32() % 2)
		}

		my := longestSubarray(arr)
		ss := make([]int, 5, 5)
		ss[0] = s01.LongestSubarray(arr)
		ss[1] = s02.LongestSubarray(arr)
		ss[2] = s03.LongestSubarray(arr)
		ss[3] = s04.LongestSubarray(arr)
		ss[4] = s05.LongestSubarray(arr)
		same := true
		for _, s := range ss {
			if s != my {
				same = false
				break
			}
		}

		if !same {
			t.Errorf(`for longestSubarray(%v), my = %d, s01 - s05 = %v, not same`, arr, my, ss)
		}
	})
}
