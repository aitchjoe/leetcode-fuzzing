package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/1139974/PythonGo-O(n)-by-stack-w-Comment

// Support function max
func Max(a, b int)int{
    
    if a > b{
        return a
    }
    
    return b
}


func longestValidParentheses(s string) int {
    
    // stack, used to record index of parenthesis
    // initialized to -1 as dummy head for valid parenthesis length computation
    stack := []int{-1}
    
    maxLength := 0
    
    // linear scan each index and character in input string s
    for curIdx, char := range s{
        
        if char == '('{
            
            // push when current char is (
            stack = append( stack, curIdx)
            
        }else{
            
            // pop when current char is )
            stack = stack[:len(stack)-1]
            
            if len(stack)==0{
                
                // stack is empty, push current index into stack
                stack = append( stack, curIdx)
            }else{
                
                // stack is non-empty, update maximal valud parentheses length
                
                maxLength = Max( maxLength, curIdx - stack[ len(stack)-1 ] )
            }
            
        }
        
    }
    
    return maxLength
    
}
