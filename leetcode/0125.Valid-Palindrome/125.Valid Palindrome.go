package leetcode

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	left, right := 0, len(s)-1
	for left < right {
		if !validCh(s[left]) {
			left++
			continue
		}
		if !validCh(s[right]) {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func validCh(ch byte) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z')
}
