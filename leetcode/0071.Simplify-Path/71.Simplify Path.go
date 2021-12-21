package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	arr := strings.Split(path, "/")
	var stack []string
	var res string
	for _, cur := range arr {
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if len(cur) > 0 && cur != "." {
			stack = append(stack, cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	res = strings.Join(stack, "/")
	return "/" + res
}
