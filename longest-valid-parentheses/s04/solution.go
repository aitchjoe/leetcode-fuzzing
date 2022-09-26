package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2304822/Simple-solution-without-stack 

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	res := 0

	nOpen, nClose := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			nOpen++
		} else {
			nClose++
		}

		if nOpen == nClose {
			res = max(res, nOpen+nClose)
		}

		if nOpen < nClose {
			nClose--
			res = max(res, nOpen+nClose)
			nOpen = 0
			nClose = 0
		}
	}

	nReverseOpen, nReverseClose := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			nReverseOpen++
		} else {
			nReverseClose++
		}

		if nReverseOpen == nReverseClose {
			res = max(res, nReverseOpen+nReverseClose)
		}

		if nReverseOpen < nReverseClose {
			nReverseClose--
			res = max(res, nReverseOpen+nReverseClose)
			nReverseOpen = 0
			nReverseClose = 0
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
