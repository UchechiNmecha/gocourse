package basics

import (
	"fmt"
	"slices"
)

func main() {
	// Declare a slice of integers
	var numbers []int
	fmt.Println(numbers)

	var numbers1 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers1)

	slice := make([]int, 5)
	fmt.Println(slice)

	// Slicing an array
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]
	fmt.Println(slice1)

	// Modifying the slice by appending elements
	slice1 = append(slice1, 6, 7, 8)
	fmt.Println(slice1)

	// copying slices
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Original Slice:", slice1)
	fmt.Println("Copied Slice:", sliceCopy)

	// nil slice
	var nilSlice []int
	fmt.Println("Nil Slice:", nilSlice)

	// Iterating over a slice
	for i, v := range slice1 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 and sliceCopy are equal")
	} else {
		fmt.Println("slice1 and sliceCopy are not equal")
	}

	// multi-dimensional slices
	var matrix [][]int = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix: ", matrix)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("Two Dimensional Slice:", twoD)

	// slice[low : high]
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sliced := s[2:5]
	fmt.Println("Sliced:", sliced)
}
