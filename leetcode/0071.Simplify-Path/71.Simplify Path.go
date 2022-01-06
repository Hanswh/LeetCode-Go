package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	ds := strings.Split(path, "/")

	var stack []string
	for _, d := range ds {
		if d == ".." {
			stack = stack[:len(stack)-1]
		} else if d != "" {
			stack = append(stack, d)
		}
	}
	res := strings.Join(stack, "/")
	return "/" + res
}
