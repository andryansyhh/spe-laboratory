package main

import (
	"math"
	"strconv"
)

func isNarcissistic(number int) bool {
	// Convert the number to a string to count its digits
	numStr := strconv.Itoa(number)
	numDigits := len(numStr)

	// Calculate the sum of the digits raised to the power of the number of digits
	var digitSum int
	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		digitSum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	// Check if the sum is equal to the original number
	return digitSum == number
}
