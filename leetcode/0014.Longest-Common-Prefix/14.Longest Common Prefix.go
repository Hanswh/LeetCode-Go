package leetcode

import "sort"

// 先按长度升序排序，再依次判断strs[0]的每一个字符是否是公共前缀的一部分
func longestCommonPrefix1(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) <= len(strs[j])
	})
	minLen := len(strs[0])
	if minLen == 0 {
		return ""
	}
	var commonPrefix []byte
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return string(commonPrefix)
			}
		}
		commonPrefix = append(commonPrefix, strs[0][i])
	}
	return string(commonPrefix)
}

// 也可以不排序，判断时加上长度
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 用LCP(s1,s2,...,sn)表示字符串s1,s2,...,sn的公共前缀
// 则LCP(s1,s2,...,sn) = LCP(LCP(LCP(s1,s2),s3),...sn)，则可以从做到右依次遍历strs并求公共前缀
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 同上使用LCP(s1,s2,...,sn)表示字符串s1,s2,...,sn的公共前缀
// LCP(s1,s2,...,sn) = LCP(LCP(s1,...,si), LCP(si+1,...,sn))
// 则除了从左向右依次进行，也可以进行分治，每次分成两部分求公共前缀，最后合并结果
func longestCommonPrefix4(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for i := 0; i < minLength; i++ {
			if lcpLeft[i] != lcpRight[i] {
				return lcpLeft[:i]
			}
		}
		return lcpLeft[:minLength]
	}
	return lcp(0, len(strs)-1)
}

// 二分搜索
// 最长公共前缀一定是最短str的一部分，其长度lenCP不会超过最短str的长度minLength
// 则可以在[0,minLength]中二分查找，找到lenCP
func longestCommonPrefix5(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		str0, count := strs[0][:length], len(strs)
		for i := 1; i < count; i++ {
			if strs[i][:length] != str0 {
				return false
			}
		}
		return true
	}
	minLength := len(strs[0])
	for _, s := range strs {
		if len(s) < minLength {
			minLength = len(s)
		}
	}
	low, high := 0, minLength
	for low < high {
		mid := (high-low+1)/2 + low
		if isCommonPrefix(mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
