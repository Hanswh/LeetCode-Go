package leetcode

import (
	"strconv"
	"strings"
)

// dfs
// 要将整个s切分得到四个sub Address，同时需要保证每个都是有效的
func restoreIpAddresses(s string) []string {
	var addresses []string

	var dfs func(index int, addr []string)
	dfs = func(index int, addr []string) {
		if len(addr) == 4 {
			// 切分成了四个
			if index == len(s) {
				// 全部遍历过，则加入结果集
				addresses = append(addresses, strings.Join(addr, "."))
			}
			return
		}
		// 从当前index开始往后找，因为要小于等于255，最多往后三个数字
		for i := 1; i <= 3; i++ {
			// 判断是否可将s[index:index+i]作为一个切分
			// 超出s长度，不能
			if index+i > len(s) {
				break
			}
			// 有前导0 || 大于255，不能
			sub, _ := strconv.Atoi(s[index : index+i])
			if i > 1 && s[index] == '0' || sub > 255 {
				break
			}

			// 切分，搜索下一个
			dfs(index+i, append(addr, s[index:index+i]))
		}
	}
	dfs(0, []string{})

	return addresses
}
