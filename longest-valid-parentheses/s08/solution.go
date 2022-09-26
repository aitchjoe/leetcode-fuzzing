package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2068392/Golang-Solution-100-faster-with-Explanation

const bar = 10000

func longestValidParentheses(s string) int {
	ret := 0
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if v == ')' {
			if len(stack) == 0 {
				continue
			}
			// If it fits in 2) the first case
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				// Try to combine the number in between
				if len(stack) > 0 && stack[len(stack)-1] > bar {
					stack[len(stack)-1] = stack[len(stack)-1] + 2
				} else {
					stack = append(stack, 2+bar)
				}
				continue
			}
			if stack[len(stack)-1] == ')' {
				continue
			}
			// if it fits in 2) second case
			if len(stack) > 1 && stack[len(stack)-2] == '(' {
				subNum := stack[len(stack)-1] - bar
				stack = stack[:len(stack)-2]
				// Try to combine the number in between
				if len(stack) > 0 && stack[len(stack)-1] > bar {
					stack[len(stack)-1] = stack[len(stack)-1] + 2 + subNum
				} else {
					stack = append(stack, 2+bar+subNum)
				}
				continue
			}
			// if it is not in the cases of 2), put it in as a partition mark
			stack = append(stack, v)
		}
	}
	// Go through all parts divided by parition mark, don't forget to substract bar.
	for _, v := range stack {
		if v > bar {
			if int(v-bar) > ret {
				ret = int(v - bar)
			}
		}
	}
	return ret
}
