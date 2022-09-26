package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2068211/Go-Stack-clear-solution

func longestValidParentheses(s string) int {
    stack, max := []int{-1}, 0
    for i, r := range s {
        switch r {
            case '(':
                stack = append(stack, i)
            case ')':
                stack = stack[:len(stack)-1]
                if len(stack) == 0 { stack = append(stack, i); continue }
                length := i-stack[len(stack)-1]
                if length > max { max = length }            
        }
    }
    return max
}
