package s03

import (
	"strings"
)

// https://leetcode.com/problems/valid-number/discuss/2575897/Share-my-very-clear-solution

func IsNumber(s string) bool {
    eIndex := strings.Index(s, "e") 
    EIndex := strings.Index(s, "E") 
    if eIndex == -1 && EIndex == -1 {
        return isDecimalNumber(s) || isInteger(s)
    } else if eIndex != -1 && EIndex != -1 {
        return false
    } else {
        var index int
        if eIndex != -1 {
            index = eIndex
        } else {
            index = EIndex
        }
        return (isDecimalNumber(s[:index]) || isInteger(s[:index])) && isInteger(s[index+1:])
    }
}

func isDecimalNumber(s string) bool {
    var hasNum bool
    var hasDot bool
    for i := 0; i < len(s); i++ {
        if i == 0 && (s[i] == '-' || s[i] == '+') {
            continue
        }
        if number(s[i]) {
            hasNum = true
            continue
        }
        if s[i] == '.' && !hasDot {
            hasDot = true
        } else {
            return false
        }
    }
    return hasNum && hasDot
}

func isInteger(s string) bool {
    var hasNum bool
    for i := 0; i < len(s); i++ {
        if i == 0 && (s[i] == '-' || s[i] == '+') {
            continue
        }
        if number(s[i]) {
            hasNum = true
            continue
        }
        return false
    }
    return hasNum
}

func number(b byte) bool {
    if b >= '0' && b <= '9' {
        return true
    }
    return false
}
