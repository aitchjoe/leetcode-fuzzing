package solution

import (
	"math/rand"
	"testing"
	"time"

	"github.com/aitchjoe/leetcode-fuzzing/valid-number/s01"
	"github.com/aitchjoe/leetcode-fuzzing/valid-number/s02"
	"github.com/aitchjoe/leetcode-fuzzing/valid-number/s03"
)

func FuzzIsNumber(f *testing.F) {
	rand.Seed(time.Now().UnixNano())
	f.Add(10)
	f.Fuzz(func(t *testing.T, n int) {
		// 1 <= s.length <= 20
		if n < 0 {
			n = -n
		}
		n = n%20 + 1

		chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-."
		s := ""
		for i := 0; i < n; i++ {
			j := int(rand.Uint32()) % len(chars)
			s += chars[j:j+1]
		}

		ss := make([]bool, 3, 3)
		ss[0] = s01.IsNumber(s)
		ss[1] = s02.IsNumber(s)
		ss[2] = s03.IsNumber(s)

		same := true
		for i := 1; i < len(ss); i++ {
			if ss[i] != ss[0] {
				same = false
				break
			}
		}

		if !same {
			t.Errorf(`for isNumber(%q), s01 - s0%d = %v, not same`, s, len(ss), ss)
		}
	})
}
