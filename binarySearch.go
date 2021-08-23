package main

import "fmt"

func main() {
	var MyArray = []int{1, 2, 3, 4, 5, 6}
	target := 2
	index := binarySearch(MyArray, target)
	if index != -1 {
		fmt.Println("The index of the target is ", index)
	} else {
		fmt.Println("The target is not found in the given array.")
	}

}

func binarySearch(myArray []int, target int) int {
	startIndex := 0
	endIndex := len(myArray) - 1
	for startIndex <= endIndex {
		mid := (startIndex + endIndex) / 2
		if target == myArray[mid] {
			return mid
		} else if target > myArray[mid] {
			startIndex = mid + 1
		} else {
			endIndex = mid - 1
		}
	}
	return -1
}
