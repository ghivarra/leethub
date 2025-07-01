func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// merge and sort first
	sorted := append(nums1, nums2...)
	sort.Ints(sorted)

	// get median based on length
	total := len(sorted)
	var result float64

	if total%2 == 0 {
		half := total / 2
		count := float64(sorted[half-1] + sorted[half])
		result = count / 2
	} else {
		indexFloat := float64(total) / 2
		index := int(indexFloat - 0.5)
		result = float64(sorted[index])
	}

	// return
	return result
}
