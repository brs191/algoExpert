/*
Write a function that takes in a non-empty array of integers that are soted in ascending order
and returns a new array of the same length with the squares of the original integers also sorted in ascending order
Note: negative values as input
*/

package main

import (
	"fmt"
)
import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	res := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		res[i] = array[i] * array[i]
	}

	sort.Ints(res)
	fmt.Printf("res %+v\n", res)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func SortedSquaredArrayOptimal(array []int) []int {
	// Write your code here.
	res := make([]int, len(array))

	large := len(array) - 1
	small := 0

	for i := len(array) - 1; i >= 0; i-- {
		largeVal := array[large]
		smallVal := array[small]

		if abs(largeVal) >= abs(smallVal) {
			res[i] = largeVal * largeVal
			large -= 1
		} else {
			res[i] = smallVal * smallVal
			small += 1
		}
	}
	fmt.Printf("optimal res %+v\n", res)
	return res
}

func main() {
	array := []int{-9, 2, 3, 4, 5, 6, 8, 9}

	SortedSquaredArray(array)
	SortedSquaredArrayOptimal(array)
}
