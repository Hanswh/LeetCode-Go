package leetcode

// ......[a,b],[c,d]...... 待插入[i,j]
// 假设i > a && i <= c
// i <= b时需要合并  leftPoint = a, rightPoint = max(b, j)
// 合并完后，还需判断[c,d]开始是不是也要合入
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	// 将p移到[c,d]处，并将之前的全部入res
	p := 0
	for p < len(intervals) && intervals[p][0] < newInterval[0] {
		res = append(res, intervals[p])
		p++
	}
	leftPoint := newInterval[0]
	rightPoint := newInterval[1]
	// [a,b]也需要合并
	if p > 0 && newInterval[0] <= intervals[p-1][1] {
		res = res[:len(res)-1]
		leftPoint = intervals[p-1][0]
		if rightPoint < intervals[p-1][1] {
			rightPoint = intervals[p-1][1]
		}
	}

	// 判断[c,d]及后面是否需要合并
	for p < len(intervals) && intervals[p][0] <= rightPoint {
		if rightPoint < intervals[p][1] {
			rightPoint = intervals[p][1]
		}
		p++
	}
	res = append(res, []int{leftPoint, rightPoint})
	for p < len(intervals) {
		res = append(res, intervals[p])
		p++
	}
	return res
}
