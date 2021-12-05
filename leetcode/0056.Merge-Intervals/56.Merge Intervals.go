package leetcode

import "sort"

// 排序 + 双指针
// 按区间左端排序区间，所有可合并的区间将连续分布
// 如果slow + 1 == fast，且intervals[fast][0] <= intervals[slow][1]，则这两个区间可以合并
// 使用一个max保存合并后的右端点值，fast++，向后寻找并判断intervals[fast][0] <= max，直到遍历完或不满足此条件，则以slow为左端点的合并完成
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	slow, fast := 0, 0
	for slow < len(intervals) {
		max := intervals[slow][0]
		for fast < len(intervals) && intervals[fast][0] <= max {
			if max < intervals[fast][1] {
				max = intervals[fast][1]
			}
			fast++
		}
		res = append(res, []int{intervals[slow][0], max})
		slow = fast
	}
	return res
}
