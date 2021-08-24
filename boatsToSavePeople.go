// you're trying to add people in safety boats
// Given an array people and an integer limit
// people's array is the array of people's weight, the i-th person has a weight people[i] and each boat can carry at most limit
// each boat carries at most 2 people at the same time, given that their weight some is at most limit.
// Return the minimum no:of boats to carry every given person
// NOTE - it is guaranteed each person can be carried by a boat.

package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []int{75}
	limit := 120
	fmt.Println("Minimum no:of boats required is: ", FindNoOfBoats(people, limit))
}

func FindNoOfBoats(peopleWeight []int, boatLimit int) int {
	sort.Ints(peopleWeight)
	left := 0
	right := len(peopleWeight) - 1
	NoOfBoats := 0
	for left <= right {
		if left == right {
			NoOfBoats++
			break
		} else {
			if peopleWeight[left]+peopleWeight[right] <= boatLimit {
				left++
				right--
				NoOfBoats++
			} else {
				right--
				NoOfBoats++
			}
		}
	}
	return NoOfBoats
}
