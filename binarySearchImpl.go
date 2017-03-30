package main

import (
	"fmt"

	"github.com/bhanushrestha/binary-search-algorithm-go-implementation/sort"
)

//Implementation of Binary Search Algorith using loop
func loopBinarySearch(sortedArray []int, startingPoint, endingPoint, target int) int {
	var midPoint int
	for startingPoint < endingPoint {
		midPoint = int((startingPoint + endingPoint) / 2)
		if sortedArray[midPoint] != target {
			if sortedArray[midPoint] < target {
				startingPoint = midPoint + 1
			} else {
				endingPoint = midPoint
			}
		} else {
			return midPoint
		}
	}
	return -1
}

//Implementation of Binary Search Algorith using recursion
func recursiveBinarySearch(sortedArray []int, startingPoint, endingPoint, target int) int {

	if startingPoint >= endingPoint {
		return -1
	}

	midPoint := int((startingPoint + endingPoint) / 2)
	if sortedArray[midPoint] != target {
		if sortedArray[midPoint] < target {
			startingPoint = midPoint + 1
		} else {
			endingPoint = midPoint
		}
		return recursiveBinarySearch(sortedArray, startingPoint, endingPoint, target)
	}
	return midPoint
}

func main() {

	//binary search requires sorted Array
	//array := []int{1, 32, 34, 45, 56, 65}
	array := []int{65, 56, 45, 34, 32, 1}
	sortingAlgorithm := sort.NewBubbleSort()
	sortingAlgorithm.SortArray(array)
	fmt.Println("Sorted Array : ", array)

	var itemToSearch int
	itemToSearch = 32
	//answer := loopBinarySearch(array, 0, len(array)-1, itemToSearch)
	answer := recursiveBinarySearch(array, 0, len(array)-1, itemToSearch)
	if answer == -1 {
		fmt.Println("VALUE not found in array")
	} else {
		fmt.Println("Value found :", array[answer], "in index ", answer, " of given array")
	}
}
