package basics

import "fmt"

func main() {

	// Declare an array of integers with a fixed size of 5
	var numbers [5]int
	fmt.Println(numbers)

	numbers[0] = 10
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println(fruits)

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray
	// copiedArray[0] = 100
	// fmt.Println(originalArray)
	// fmt.Println(copiedArray)

	// for i := 0; i < len(originalArray); i++ {
	// 	fmt.Println("Element at index", i, "is", originalArray[i])
	// }

	for index, value := range fruits {
		fmt.Println("Element at index", index, "is", value)
	}

	// Blank identifier to ignore the index
	for _, value := range fruits {
		fmt.Printf("fruits - %s\n", value)
	}

	a, b := someFunction()
	fmt.Println(a, b)

	_, value2 := someFunction()
	fmt.Println(value2)

	// Comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("Are array1 and array2 equal?", array1 == array2)

	var array3 [3]int = [3]int{4, 5, 6}
	fmt.Println("Are array1 and array3 equal?", array1 == array3)

	// multi-dimensional arrays
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix: ", matrix)

	var originalArray [3]int = [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100
	fmt.Println(originalArray)
	fmt.Println(copiedArray)
}

func someFunction() (int, string) {
	return 42, "Hello"
}
