package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/1316357/Golang-stack-8ms3.8mb-with-a-character-validation

func longestValidParentheses(s string) int {
    // initialize the stack
    stack := make([]int, 0)
    // valid characters
	mappings := make(map[byte]bool,2)
    mappings[')'] = true
    mappings['('] = true
    maxans := 0
    
	// insert 1 element into stack
    stack = append(stack, -1)
    
    for i := 0; i<len(s);i++ {
        // read current char
        char := s[i]
        
        // is it an open or closing char?
        if mappings[char] == true {
            if char == '(' {
                // append index
                stack = append(stack, i)
            } else {
                // pop()
                stack = stack[:len(stack)-1]
                
                // if the stack is empty, push()
                if len(stack) == 0 {
                     stack = append(stack, i)
                } else {
                    // otherwise read the top
                    // and check the max
                    top := stack[len(stack)-1]
                    maxans = max(maxans, i- top)
                }
            }
        }
    }
    
    return maxans
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}
