package s02

import (
	"strings"
)

// https://leetcode.com/problems/valid-number/discuss/23915/Share-my-golang-solution

import (
    "unicode"
    "unicode/utf8"
)

func IsNumber(s string) bool {
    t := strings.TrimSpace(s)
    i, n := 0, len(t)

    if i < n && (string(t[i]) == "+" || string(t[i]) == "-") {
	    i++
    }

    isNumeric := false
    for i < n && isDigit(t[i]) {
	    i++
	    isNumeric = true
    }

    if i < n && string(t[i]) == "." {
	    i++
	    for i < n && isDigit(t[i]) {
		    i++
		    isNumeric = true
	    }
    }

    if isNumeric && i < n && string(t[i]) == "e" {
	    i++
	    isNumeric = false
	    if i < n && (string(t[i]) == "+" || string(t[i]) == "-") {
		    i++
	    }
	    for i < n && isDigit(t[i]) {
		    i++
		    isNumeric = true
	    }
    }

    return isNumeric && i == n
}

func isDigit(d byte) bool {
    digit, _ := utf8.DecodeRune([]byte(string(d)))
    return unicode.IsDigit(digit)
}
