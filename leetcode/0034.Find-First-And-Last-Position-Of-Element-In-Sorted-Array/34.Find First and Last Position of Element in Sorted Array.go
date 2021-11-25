package leetcode

// 解法一: 遍历
func searchRange1(nums []int, target int) []int {
	i, j := 0, 0
	for i = 0; i < len(nums); i++ {
		if nums[i] == target {
			break
		}
	}
	for j = i; j < len(nums)-1; j++ {
		if nums[j+1] != target {
			break
		}
	}
	if i == len(nums) {
		i, j = -1, -1
	}
	return []int{i, j}
}

// 解法二: 二分加遍历
func searchRange(nums []int, target int) []int {
	done := false

	low, high := 0, len(nums)-1

	// {5,7,7,8,8,10,11} 8
	for low <= high && !done {
		mid := low + (high-low)>>1

		if nums[mid] == target { // mid落到了target序列上
			for nums[low] != target {
				low++
			}
			for nums[high] != target {
				high--
			}
			done = true
		} else if nums[mid] > target { // mid落到了target序列右方
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !done {
		low, high = -1, -1
	}
	return []int{low, high}
}

// 解法三: 二分
func searchRange3(nums []int, target int) []int {
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

// 二分查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == 0) || (nums[mid-1] != target) { // 找到第一个与 target 相等的元素
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 找到最后一个与 target 相等的元素
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

// 二分查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// 二分查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if (mid == len(nums)-1) || (nums[mid+1] > target) { // 找到最后一个小于等于 target 的元素
				return mid
			}
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
