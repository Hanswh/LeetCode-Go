package leetcode

// 笨方法，从左到右对每个i，判断s[i:i+subL]是否满足要求，subL=len(words)*len(words[0])
// 没有利用好单词长度一致的规律去使用滑动窗口减少词频统计的次数
func findSubstring1(s string, words []string) []int {
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}

	subL, wordL := len(words)*len(words[0]), len(words[0])
	var res []int
	for start := 0; start < len(s)-subL+1; start++ {
		left := start
		window := make(map[string]int)
		for k := range count {
			window[k] = count[k]
		}
		for left < start+subL {
			right := left + wordL
			if c := window[s[left:right]]; c > 0 {
				window[s[left:right]]--
			} else {
				break
			}
			left = right
		}
		if left == start+subL {
			res = append(res, start)
		}
	}
	return res
}

// 好办法，从每个i < wordLen开始, 使用滑动窗口(每次滑动距离wordLen)依次检测s[i:i+subL]是否满足要求
func findSubstring2(s string, words []string) []int {
	var result []int // 最终结果
	// 特判
	if len(s) == 0 || len(words) == 0 {
		return result
	}
	// 统计words中每个单词出现的次数
	need := make(map[string]int)
	for _, word := range words {
		need[word]++
	}
	wordLen, subLen, ls := len(words[0]), len(words[0])*len(words), len(s)
	for start := 0; start < wordLen; start++ {
		// 窗口左边，窗口右边，词频匹配的单词数
		left, right, match := start, start, 0
		// 窗口中每个单词出现的次数
		window := make(map[string]int)
		for right+wordLen <= ls {
			rightWord := s[right : right+wordLen] // 当前新加入的单词
			right += wordLen
			if need[rightWord] > 0 {
				window[rightWord]++
				if window[rightWord] == need[rightWord] {
					match++
				}
			}
			// 如果满足了长度，判断是否满足词频
			if right-left == subLen {
				// 词频也匹配，加入结果
				if match == len(need) {
					result = append(result, left)
				}
				leftWord := s[left : left+wordLen] // 当前新移出的单词
				left += wordLen
				if need[leftWord] > 0 {
					if window[leftWord] == need[leftWord] {
						match--
					}
					window[leftWord]--
				}
			}
		}
	}
	return result
}
