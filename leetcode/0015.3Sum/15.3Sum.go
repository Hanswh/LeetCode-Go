package leetcode

import "sort"

func threeSum1(nums []int) [][]int {
	var res [][]int
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	var uniqNums []int
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}

// 排序 + 双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	index, addSum, length, start, end := 0, 0, len(nums), 0, 0

	for index = 1; index < length-1; index++ {
		start, end = 0, length-1

		if index > 0 && nums[index] == nums[index-1] {
			start = index - 1
		}

		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}

			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			addSum = nums[start] + nums[index] + nums[end]
			if addSum == 0 {
				res = append(res, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addSum < 0 {
				start++
			} else {
				end--
			}
		}
	}
	return res
}

// 排序 + 双指针
func threeSum3(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
