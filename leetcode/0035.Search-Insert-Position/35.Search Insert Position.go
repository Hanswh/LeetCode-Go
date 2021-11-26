package leetcode

// 不满足题意，O(n) 从左往右遍历找到第一个满足nums[i] >= target的i
func searchInsert1(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); i++ {
		if nums[i] >= target {
			break
		}
	}
	return i
}

// 二分查找，找从左往右第一个满足nums[i] >= target的i
func searchInsert2(nums []int, target int) int {
	// [1,3] 2
	// low = 0, high = 1 expect 1
	// (从左往右第一个满足nums[i] >= target的i) 等价于 (满足nums[i-1] < target <= nums[i]的i)
	// 初始解空间[0, len(nums)]，我们把解空间使用[low, high+1]表示，一步步逼近直至只剩一个元素即low==high+1
	low, high := 0, len(nums)-1
	for low <= high { // 退出条件low>high实质上是low==high+1，此时[low, high+1]变为[low, low] 或 [high+1, high+1]，唯一确定了i
		mid := low + (high-low)>>1
		if nums[mid] < target {
			// nums[mid] 小于 target，则我们取low = mid + 1
			// 因为要取的nums[i] >= target，而nums[mid]小于target，则i在[mid+1, high+1]里，也即更新后的[low, high+1]中
			low = mid + 1
		} else {
			// nums[mid] 大于等于target，则我们取high = mid - 1
			// 因为要取的nums[i] >= target，而nums[mid]大于等于target，则i在[low, mid]里，也即更新后的[low, high+1]中
			high = mid - 1
		}
	}
	return low
}

// 延申
// 二分查找，找从左往右最后一个小于等于target的nums[i]
func searchInsert3(nums []int, target int) int {
	// （从左往右最后一个小于target的nums[i]） 等价于 （满足nums[i] <= target < nums[i+1]的i)
	// 初始解空间[-1, len(nums)-1]，我们把解空间用[low-1, high]表示，一步步逼近直至只剩一个元素即low-1==high
	low, high := 0, len(nums)-1
	for low <= high { // 循环退出条件low>high实质上是low-1==high，此时[low-1,high]变为[high,high] 或 [low-1,low-1]，唯一确定了i
		mid := low + (high-low)>>1
		if nums[mid] <= target {
			// nums[mid] 小于等于target，则low = mid + 1
			// 要取的nums[i] <= target，而nums[mid]小于等于target，则i在[mid, high]里，也即更新后的[low-1,high]中
			low = mid + 1
		} else {
			// nums[mid] 大于target，则取high = mid - 1
			// 要取的nums[i] > target，而nums[mid]大于target，则i在[low-1, mid-1]里，也即更新后的[low-1,high]中
			high = mid - 1
		}
	}
	return high
}
