package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2069886/go-stack-solution

func longestValidParentheses(s string) int {
    var parentheses []int
    parentheses = append(parentheses, -1)
    result := 0
    for i := 0; i < len(s); i++ {
        if s[i] == 40 {
            parentheses = append(parentheses, i)
            continue
        }
        if len(parentheses) == 0 {
            continue
        }
        parentheses = parentheses[:len(parentheses)-1]
        if len(parentheses) == 0 {
            parentheses = append(parentheses, i)
        } else {
            topIndex := len(parentheses) - 1
            if result < i - parentheses[topIndex] {
                result = i - parentheses[topIndex]
            }
        }
    }
    return result
}
