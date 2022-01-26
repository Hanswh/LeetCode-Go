package leetcode

// 笨方法，边转后缀边计算
func calculate(s string) int {
	var ops []string // 运算符栈
	var nums []int   // 操作数栈

	// handle 根据传入的参数对两个栈进行操作，参数中的待处理子串cur都是双目运算符和操作数
	// 算符优先算法，根据优先级处理
	// 当前读入的是操作数时直接入操作数栈
	// 当前读入的是双目运算符时，如果优先级高于栈顶则入栈，低于栈顶则计算（操作数栈顶再前一个 运算符栈顶 操作数栈顶）至高于栈顶入栈
	var handle func(cur string, isNum bool)
	handle = func(cur string, isNum bool) {
		if isNum {
			nums = append(nums, atoi(cur))
		} else {
			switch cur {
			case "+", "-":
				if ops[len(ops)-1] == "(" {
					ops = append(ops, cur)
				} else {
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
					ops[len(ops)-1] = cur
				}
			case "*", "/":
				if ops[len(ops)-1] == "(" || ops[len(ops)-1] == "+" || ops[len(ops)-1] == "-" {
					ops = append(ops, cur)
				} else {
					switch ops[len(ops)-1] {
					case "*":
						nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
					case "/":
						nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
					}
					nums = nums[:len(nums)-1]
					ops[len(ops)-1] = cur
				}
			case "(":
				ops = append(ops, cur)
			case ")":
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
			}
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

// 因为只这题只有加减号，可以简化
// 思路是让每个数字前面的运算符都变成展开括号后的值，运算后加到结果上
// 我们用栈保存每一层括号前面的符号，即(1+1)还是-(1+1)
func calculate2(s string) (ans int) {
	ops := []int{1}
	sign := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return
}
