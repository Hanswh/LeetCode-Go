package leetcode

// 从小到大分解num
func intToRoman1(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}

// 硬编码数字
var (
	thousands = []string{"", "M", "MM", "MMM"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func intToRoman2(num int) string {
	return thousands[num/1000] + hundreds[num%1000/100] + tens[num%100/10] + ones[num%10]
}
