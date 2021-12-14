package leetcode

// 蛮力法
func longestValidParentheses0(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if isValidParentheses(s[i:j]) && maxLen < j-i {
				maxLen = j - i
			}
		}
	}
	return maxLen
}

func isValidParentheses(s string) bool {
	count := 0
	for i := range s {
		if s[i] == '(' {
			count++
		} else if count > 0 {
			count--
		} else {
			return false
		}
	}
	return count == 0
}

// 解法一 栈
// 栈中元素表示下标，在栈中我们维护一个“上一个无效的括号序列的最右端下标”（栈顶）。初始时，这个值是-1，且无法匹配，可以看作一个右括号
// 当遇到左括号时将其入栈，待匹配，“上一个无效的括号序列的最右端下标”变成刚刚入栈的左括号
// 当遇到右括号时将栈顶元素出栈，有两种情况：出栈的是一个左括号；出栈的是“上一个无效的括号序列的最右端下标”；
// 	如果栈中有左括号，右括号匹配成功，这里左括号出栈后栈不空，因为栈底始终有个无法匹配的
//		要做的操作是更新一下最大值res = max(res, i和当前栈顶值的差）。
// 	如果栈中没有左括号，则遇到了无法匹配的右括号，其位于栈底，出栈后栈空
//		要做的操作是再将当前右括号下标入栈，作为新的无法匹配的“上一个无效的括号序列的最右端下标”
func longestValidParentheses1(s string) int {
	var stack []int
	res := 0
	stack = append(stack, -1)
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

// 解法二 双指针
func longestValidParentheses2(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
