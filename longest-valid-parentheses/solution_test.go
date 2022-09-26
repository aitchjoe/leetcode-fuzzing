package solution

import (
	"math/rand"
	"testing"
	"time"

	s01 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s01"
	s02 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s02"
	s03 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s03"
	s04 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s04"
	s05 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s05"
	s06 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s06"
	s07 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s07"
	s08 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s08"
	s09 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s09"
	s10 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s10"
	s11 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s11"
	s12 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s12"
	s13 "github.com/aitchjoe/leetcode-fuzzing/longest-valid-parentheses/s13"
)

func Fuzz(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(303)
	f.Fuzz(func(t *testing.T, n int) {
		// Constraints: 0 <= s.length <= 3 * 10^4
		if n < 0 {
			n = -n
		}
		n = n%30001

		b := make([]byte, n, n)
		for i := 0; i < n; i++ {
			j := int(rand.Uint32()) % 2
			// Constraints: 0 <= s.length <= 3 * 104
			if j == 0 {
				b[i] = '('
			} else {
				b[i] = ')'
			}
		}
		s := string(b)

		r := make([]int, 13, 13)
		r[0] = s01.LongestValidParentheses(s)
		r[1] = s02.LongestValidParentheses(s)
		r[2] = s03.LongestValidParentheses(s)
		r[3] = s04.LongestValidParentheses(s)
		r[4] = s05.LongestValidParentheses(s)
		r[5] = s06.LongestValidParentheses(s)
		r[6] = s07.LongestValidParentheses(s)
		r[7] = s08.LongestValidParentheses(s)
		r[8] = s09.LongestValidParentheses(s)
		r[9] = s10.LongestValidParentheses(s)
		r[10] = s11.LongestValidParentheses(s)
		r[11] = s12.LongestValidParentheses(s)
		r[12] = s13.LongestValidParentheses(s)

		same := true
		for i := 1; i < len(r); i++ {
			if r[i] != r[0] {
				same = false
				break
			}
		}

		if !same {
			t.Errorf(`for longestValidParentheses(%q), s01 - s13 = %v, not same`, s, r)
		}
	})
}
