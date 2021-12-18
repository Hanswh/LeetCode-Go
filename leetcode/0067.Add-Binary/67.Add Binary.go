package leetcode

func addBinary(a string, b string) string {
	n := len(a)
	if len(a) < len(b) {
		n = len(b)
	}
	res := ""
	var sum, carry, numA, numB byte
	for i := 0; i < n; i++ {
		numA = 0
		numB = 0
		if i < len(a) {
			numA = a[len(a)-i-1] - '0'
		}
		if i < len(b) {
			numB = b[len(b)-i-1] - '0'
		}
		sum = numA + numB + carry
		carry = sum / 2
		sum %= 2
		res = string(sum+'0') + res
	}
	if carry == 1 {
		res = "1" + res
	}
	return res
}
