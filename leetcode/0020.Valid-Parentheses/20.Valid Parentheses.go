package leetcode

var parentheses = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == parentheses[s[i]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
