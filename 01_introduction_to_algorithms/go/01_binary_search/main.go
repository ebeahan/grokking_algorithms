package main

import (
	"fmt"
)

func main() {
	myList := []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(myList, 3))

	// Returns nil
	fmt.Println(binarySearch(myList, -1))
}

func binarySearch(list []int, item int) bool {
	// low and high keep track of which part of the list you'll search in
	low := 0
	high := len(list) - 1

	// While you haven't narrowed it down to one element...
	for low <= high {
		//... check the middle element
		mid := (low + high) / 2
		guess := list[mid]
		// Found the item.
		if guess == item {
			return true
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// Item doesn't exist
	return false
}
