package leetcode

var roman = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

// MCM 1900
func romanToInt(s string) int {
	res, lastInt, num := 0, 0, 0
	for i := 0; i < len(s); i++ {
		num = roman[s[len(s)-(i+1):len(s)-i]]
		if lastInt > num {
			res -= num
		} else {
			res += num
		}
		lastInt = num
	}
	return res
}
