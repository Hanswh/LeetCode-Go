package leetcode

import "sort"

// 排序，使用字典顺序的单词作为键
func groupAnagrams1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		tmp := []byte(str)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		sortedStr := string(tmp)
		m[sortedStr] = append(m[sortedStr], str)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 计数，使用各个字母出现的次数作为键
func groupAnagrams2(strs []string) [][]string {
	m := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	ans := make([][]string, 0, len(m))
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
