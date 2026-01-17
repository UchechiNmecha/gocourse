package basics

import "fmt"

func main() {
	// add := add(1, 2)
	// fmt.Println("The sum is: ", add)

	greet := func() {
		fmt.Println("This is an anonymous function.")
	}
	greet()

	operation := add
	result := operation(3, 4)
	fmt.Println("The result of the operation is: ", result)

	// passing function as parameter
	sum := applyOperation(2, 5, add)
	fmt.Println("2 + 5: ", sum)

	// returning function
	multipleBy2 := createMultiplier(2)
	fmt.Println("5 multiplied by 2 is: ", multipleBy2(5))

}

func add(a int, b int) int {
	return a + b
}

// A function that takes another function as a parameter
func applyOperation(x int, y int, operation func(a int, b int) int) int {
	return operation(x, y)
}

// A function that returns another function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
