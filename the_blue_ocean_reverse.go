package main

func blueOcean(list1, list2 []int) []int {
	// Create a map to store values from list2 for efficient lookup
	valuesToRemove := make(map[int]bool)
	for _, val := range list2 {
		valuesToRemove[val] = true
	}

	// Create a result slice to store values from list1 that are not in list2
	result := []int{}
	for _, val := range list1 {
		if !valuesToRemove[val] {
			result = append(result, val)
		}
	}

	return result
}
