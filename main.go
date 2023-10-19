package main

import (
	"fmt"
)

func main() {
	// number 1
	fmt.Println("NUMBER 1", isNarcissistic(153)) // true
	fmt.Println("NUMBER 1", isNarcissistic(111)) // false

	//number 2
	arr1 := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	fmt.Println("NUMBER 2", findOutlier(arr1))

	arr2 := []int{160, 3, 1719, 19, 11, 13, -21}
	fmt.Println("NUMBER 2", findOutlier(arr2))

	arr3 := []int{11, 13, 15, 19, 9, 13, -21}
	fmt.Println("NUMBER 2", findOutlier(arr3))

	//number 3
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	needle := "blue"
	index := findNeedle(haystack, needle)
	fmt.Println("NUMBER 3", index)

	//number 4
	list1 := []int{1, 2, 3, 4, 6, 10}
	list2 := []int{1}
	result := blueOcean(list1, list2)
	fmt.Println("NUMBER 4", result)

	list1 = []int{1, 5, 5, 5, 5, 3}
	list2 = []int{5}
	result = blueOcean(list1, list2)
	fmt.Println("NUMBER 4", result)
}
