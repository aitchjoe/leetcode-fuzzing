package solution

// Public access for fuzz testing.
func LongestValidParentheses(s string) int {
	return longestValidParentheses(s)
}


// https://leetcode.com/problems/longest-valid-parentheses/discuss/2073046/Golang-(Go)-solution-Beats-100

type Stack []int

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str int) {
	*s = append(*s, str)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return 0
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element
	}
}

func longestValidParentheses(s string) int {
	res := 0

	var stack Stack
	stack.Push(-1)

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			stack.Push(i)
		} else {
			stack.Pop()

			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				num := stack.Peek()
				res = max(res, i-num)
			}
		}
	}

	return res
}
