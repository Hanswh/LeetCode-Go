package leetcode

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	rankDesc := [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}

	type pair struct {
		s int // score
		i int // index
	}
	arr := make([]pair, len(score))
	for i, s := range score {
		arr[i] = pair{s, i}
	}
	// 按得分降序排列
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].s > arr[j].s
	})

	res := make([]string, len(score))
	for i, p := range arr {
		if i < 3 {
			res[p.i] = rankDesc[i]
		} else {
			res[p.i] = strconv.Itoa(i + 1)
		}
	}

	return res
}
