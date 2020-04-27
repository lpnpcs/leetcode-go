package src

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var n int = len(nums1)
	var m int = len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	return float64(getKth(nums1, 0, n-1, nums2, 0, m-1, left) + getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5
}

func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) int {
	len1 := end1 - start1 + 1
	len2 := end2 - start2 + 1
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return getMin(nums1[start1], nums2[start2])
	}

	i := start1 + getMin(len1, k/2) - 1
	j := start2 + getMin(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}

func getMin(var1 int, var2 int) int {
	if var1 > var2 {
		return var2
	} else {
		return var1
	}
}
