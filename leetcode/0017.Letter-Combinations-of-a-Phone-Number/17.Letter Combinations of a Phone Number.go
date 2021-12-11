package leetcode

var letters = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var res []string

	var dfs func(subRes string, index int)
	dfs = func(subRes string, index int) {
		if index == len(digits) {
			if len(subRes) > 0 {
				res = append(res, subRes)
			}
			return
		}
		i := digits[index] - '2'
		for _, letter := range letters[i] {
			dfs(subRes+letter, index+1)
		}
	}
	dfs("", 0)
	return res
}
