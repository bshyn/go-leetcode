package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeSortedArrays(nums1, nums2)

	if len(merged) % 2 == 0 {
		return float64(merged[len(merged)/2] + merged[len(merged)/2-1]) / 2
	} else {
		return float64(merged[len(merged)/2])
	}
}
