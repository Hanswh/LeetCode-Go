package leetcode

import "math"

func myAtoi(s string) int {
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}

	sign := 1
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	var res int64
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int64(s[i]-'0')
		if tmp := res * int64(sign); tmp > math.MaxInt32 {
			return math.MaxInt32
		} else if tmp < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}
	return sign * int(res)
}
