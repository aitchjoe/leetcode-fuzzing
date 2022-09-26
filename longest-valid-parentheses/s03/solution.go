package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2462220/Go-solution 

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {

	stack := []int{-1}
	res := 0

	for i, c := range s {

		if c == '(' {
			// push
			stack = append(stack, i)

		} else {
			// pop
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				// push
				stack = append(stack, i)

			} else {
				res = max(res, i-stack[len(stack)-1])

			}

		}

	}

	return res
}
