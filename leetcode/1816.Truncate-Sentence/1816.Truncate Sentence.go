package leetcode

func truncateSentence(s string, k int) string {
	i := 0
	for k > 0 {
		i++
		if i == len(s) || s[i] == ' ' {
			k--
		}
	}
	return s[:i]
}
