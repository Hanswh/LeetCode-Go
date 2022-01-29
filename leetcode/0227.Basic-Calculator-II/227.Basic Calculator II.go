package leetcode

// 边转后缀边计算，额外考虑了括号，和题224方法一代码一样
func calculate1(s string) int {
	var ops []string // 运算符栈
	var nums []int   // 操作数栈

	// handle 根据传入的参数对两个栈进行操作，参数中的待处理子串cur都是双目运算符和操作数
	// 算符优先算法，根据优先级处理
	// 当前读入的是操作数时直接入操作数栈
	// 当前读入的是双目运算符时，如果优先级高于栈顶则入栈，低于栈顶则计算（操作数栈顶再前一个 运算符栈顶 操作数栈顶）至高于栈顶入栈
	// 当前读入的是左括号时直接入栈，右括号时计算（操作数栈顶再前一个 运算符栈顶 操作数栈顶）至栈顶是左括号，将其出栈
	var handle func(cur string, isNum bool)
	handle = func(cur string, isNum bool) {
		if isNum {
			nums = append(nums, atoi(cur))
		} else if cur == "(" {
			ops = append(ops, cur)
		} else if cur == ")" {
			for ops[len(ops)-1] != "(" {
				switch ops[len(ops)-1] {
				case "+":
					nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
				case "-":
					nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
				case "*":
					nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
				case "/":
					nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
				}
				nums = nums[:len(nums)-1]
				ops = ops[:len(ops)-1]
			}
			ops = ops[:len(ops)-1]
		} else {
			for higher(ops[len(ops)-1], cur) {
				switch ops[len(ops)-1] {
				case "+":
					nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
				case "-":
					nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
				case "*":
					nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
				case "/":
					nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
				}
				nums = nums[:len(nums)-1]
				ops = ops[:len(ops)-1]
			}
			ops = append(ops, cur)
		}
	}

	handle("(", false)
	// 双指针法寻找下一个待处理子串
	// 子串分为三类：操作数、双目运算符、单目运算符
	i, j := 0, 0
	for i < len(s) {
		j = i + 1
		if s[i] == ' ' {
			i = j
			continue
		}
		if (i == 0 || s[i-1] == '(') && s[i] == '-' {
			// 遇到单目'-'，这里采用的处理方式是模拟一个零减的操作，把单目'-'转化为双目'-'
			handle("0", true)
			handle("-", false)
			i = j
			continue
		}
		isNum := false
		if s[i] >= '0' && s[i] <= '9' {
			// 遇到数字
			isNum = true
			for j < len(s) && s[j] >= '0' && s[j] <= '9' {
				j++
			}
		}
		handle(s[i:j], isNum)
		i = j
	}
	handle(")", false)

	return nums[0]
}

func higher(op1, op2 string) bool {
	switch op1 {
	case "(":
		return false
	case "+", "-":
		return op2 != "*" && op2 != "/"
	case "*", "/":
		return true
	}
	return false
}

// atoi 把数字字符串转化为对应值的int
func atoi(s string) int {
	flag := 1 // 是否是负数
	if s[0] == '-' {
		flag = -1
		s = s[1:]
	}

	n := 0
	for _, c := range s {
		n = n*10 + int(c-'0')
	}
	return n * flag
}

// 先算乘除，再算加减
func calculate2(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	// 先计算乘除
	//
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	// 加减
	for _, v := range stack {
		ans += v
	}
	return
}
