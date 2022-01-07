package leetcode

// 普通插入，有大量移动操作
func merge1(nums1 []int, m int, nums2 []int, n int) {
	index := 0 // 插入位置
	for i := range nums2 {
		// 找插入位置
		for ; index < m+i && nums1[index] <= nums2[i]; index++ {
		}
		// 后移
		copy(nums1[index+1:m+i+1], nums1[index:m+i])
		// 插入
		nums1[index] = nums2[i]
	}
}

// 普通选择，额外空间
func merge2(nums1 []int, m int, nums2 []int, n int) {
	nums := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	nums = append(nums, nums1[i:m]...)
	nums = append(nums, nums2[j:n]...)
	copy(nums1, nums)
}

// 优化选择
// 从nums1[m+n-1]开始，挑选最大的，这样不需要额外空间，也可以只复制一次
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for index := m + n - 1; m > 0 && n > 0; index-- {
		if nums1[m-1] > nums2[n-1] {
			nums1[index] = nums1[m-1]
			m--
		} else {
			nums1[index] = nums2[n-1]
			n--
		}
	}
	// 如果是m>0则不用处理，如果是n>0则依次复制
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}
}
