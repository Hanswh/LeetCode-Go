package leetcode

func multiply(num1 string, num2 string) string {
	product := make([]int, len(num1)+len(num2)) // 这里使用int，因为过程中会溢出
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			product[i+j+1] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	// 进位，得出结果
	res := make([]byte, len(product))
	for i := len(product) - 1; i > 0; i-- {
		product[i-1] += product[i] / 10
		res[i] = byte(product[i]%10) + '0'
	}
	res[0] = byte(product[0]) + '0'
	// 去除前缀0
	index := 0
	for index < len(res)-1 {
		if res[index] != '0' {
			break
		}
		index++
	}
	return string(res[index:])
}
