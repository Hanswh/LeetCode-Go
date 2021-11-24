package leetcode

import (
	"sort"
)

func fourSum1(nums []int, target int) [][]int {
	counter, uniqNums, res := make(map[int]int), make([]int, 0, 0), make([][]int, 0, 0)

	for _, num := range nums {
		counter[num]++
	}

	for key := range counter {
		uniqNums = append(uniqNums, key)
	}

	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*4 == target) && (counter[uniqNums[i]] > 3) {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*3+uniqNums[j] == target) && counter[uniqNums[i]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[i]*2+uniqNums[j]*2 == target) && counter[uniqNums[i]] > 1 && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			if (uniqNums[i]+uniqNums[j]*3 == target) && counter[uniqNums[j]] > 2 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[j]})
			}
			for k := j + 1; k < len(uniqNums); k++ {
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[i]+uniqNums[j]*2+uniqNums[k] == target) && counter[uniqNums[j]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j], uniqNums[k]})
				}
				if (uniqNums[i]+uniqNums[j]+uniqNums[k]*2 == target) && counter[uniqNums[k]] > 1 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], uniqNums[k]})
				}

				need := target - uniqNums[i] - uniqNums[j] - uniqNums[k]
				if need > uniqNums[k] && counter[need] > 0 {
					res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[k], need})
				}
			}
		}
	}
	return res
}

// 解法二 双指针
func fourSum2(nums []int, target int) [][]int {
	sort.Ints(nums)

	res, length := make([][]int, 0, 0), len(nums)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, length-1
			for k < l {
				addSum := nums[i] + nums[j] + nums[k] + nums[l]
				if addSum == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					for k++; k < l && nums[k] == nums[k-1]; k++ {
					}
					for l--; k < l && nums[l] == nums[l+1]; l-- {
					}
				} else if addSum > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return res
}

// 解法三 kSum
func fourSum3(nums []int, target int) [][]int {
	res, cur := make([][]int, 0), make([]int, 0)
	sort.Ints(nums)
	kSum(nums, 0, len(nums)-1, target, 4, cur, &res)
	return res
}

func kSum(nums []int, left, right int, target int, k int, cur []int, res *[][]int) {
	if right-left+1 < k || k < 2 || target < nums[left]*k || target > nums[right]*k {
		return
	}
	if k == 2 {
		// 2 sum
		twoSum(nums, left, right, target, cur, res)
	} else {
		for i := left; i < len(nums); i++ {
			if i == left || (i > left && nums[i-1] != nums[i]) {
				next := make([]int, len(cur))
				copy(next, cur)
				next = append(next, nums[i])
				kSum(nums, i+1, len(nums)-1, target-nums[i], k-1, next, res)
			}
		}
	}

}

func twoSum(nums []int, left, right int, target int, cur []int, res *[][]int) {
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			cur = append(cur, nums[left], nums[right])
			temp := make([]int, len(cur))
			copy(temp, cur)
			*res = append(*res, temp)
			// reset cur to previous state
			cur = cur[:len(cur)-2]
			left++
			right--
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
}

func fourSum4(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
