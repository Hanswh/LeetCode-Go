package leetcode

func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for right >= 0 && s[right] == ' ' {
		right--
	}
	left := right - 1
	for left >= 0 && s[left] != ' ' {
		left--
	}
	return right - left
}
