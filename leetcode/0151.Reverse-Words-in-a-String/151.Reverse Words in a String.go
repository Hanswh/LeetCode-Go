package leetcode

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	res := ""
	for _, word := range words {
		if word != "" {
			res = word + " " + res
		}
	}
	return res[:len(res)-1]
}
