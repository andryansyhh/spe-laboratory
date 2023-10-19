package main

func findOutlier(arr []int) int {
	evenCount, oddCount := 0, 0
	var even, odd int

	for _, num := range arr {
		if num%2 == 0 {
			evenCount++
			even = num
		} else {
			oddCount++
			odd = num
		}

		// If we have identified both even and odd numbers, we can exit the loop early.
		if evenCount > 0 && oddCount > 0 {
			break
		}
	}

	// Determine the outlier based on counts.
	if evenCount == 1 {
		return even
	} else if oddCount == 1 {
		return odd
	}

	// If no outlier is found, return false.
	return 0
}
