package leetcode

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, s := range tokens {
		switch s {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, toInt(s))
		}
	}
	return stack[0]
}

func toInt(s string) int {
	flag := 1
	if s[0] == '-' {
		flag = -1
		s = s[1:]
	}
	num := 0
	for _, d := range s {
		num = num*10 + int(d-'0')
	}
	return num * flag
}
