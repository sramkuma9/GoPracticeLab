// Given an array of integers write a function to move all zero's to the ned while maintaining the relative order of the other elements.

package main

import (
	"fmt"
)

func moveZeros(mySlice []int) []int {
	zeroCount := 0
	for i, _ := range mySlice {
		if mySlice[i] != 0 {
			mySlice[zeroCount] = mySlice[i]
			zeroCount++
		}
	}
	for zeroCount < len(mySlice) {
		mySlice[zeroCount] = 0
		zeroCount++
	}
	return mySlice
}

func main() {
	mySlice := []int{0, 6, 0, 8, 0}
	fmt.Println("Given Array is: ", mySlice)
	resultSlice := moveZeros(mySlice)
	fmt.Println("Result Array is: ", resultSlice)
}
