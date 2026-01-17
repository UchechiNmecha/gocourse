package basics

import "fmt"

func main() {

	// valid input
	process(10)

	// invalid input to demonstrate panic
	process(-5)
}

func process(input int) {
	defer fmt.Println("Defer 1")
	defer fmt.Println("Defer 2")

	if input < 0 {
		fmt.Println("This is before panic")
		panic("Input must be a non-negative number")
	}
	fmt.Println("Processing input:", input)
}
