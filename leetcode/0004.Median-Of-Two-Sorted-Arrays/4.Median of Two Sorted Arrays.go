package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low, high := 0, len(nums1)
	k := (len(nums1) + len(nums2) + 1) >> 1
	mid1, mid2 := 0, 0
	for low <= high {
		mid1 = low + (high-low)>>1
		mid2 = k - mid1

		if mid1 > 0 && nums1[mid1-1] > nums2[mid2] {
			high = mid1 - 1
		} else if mid1 != len(nums1) && nums2[mid2-1] > nums1[mid1] {
			low = mid1 + 1
		} else {
			break
		}
	}
	midLeft, midRight := 0, 0
	if mid1 == 0 {
		midLeft = nums2[mid2-1]
	} else if mid2 == 0 {
		midLeft = nums1[mid1-1]
	} else {
		midLeft = max(nums1[mid1-1], nums2[mid2-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if mid1 == len(nums1) {
		midRight = nums2[mid2]
	} else if mid2 == len(nums2) {
		midRight = nums1[mid1]
	} else {
		midRight = min(nums1[mid1], nums2[mid2])
	}
	return float64(midLeft+midRight) / 2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
