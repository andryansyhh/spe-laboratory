package main

func findNeedle(haystack []string, needle string) int {
	for i, value := range haystack {
		if value == needle {
			return i
		}
	}

	// Return -1 if the needle is not found in the haystack.
	return -1
}
