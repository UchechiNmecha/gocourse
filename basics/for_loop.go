package basics

import "fmt"

func main() {
	// Simple for loop from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over a collection
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Number: %d, Value: %d\n", index, value)
	}

	for i := 1; i <=  10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd Numbers: ", i)
		if i == 5 {
			break
		}
	}


	// Nested for loops to print a pattern
	rows := 10
	for i:= 1; i <= rows; i++ {

		// first tree
		for j := 1; j <= rows - i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i - 1; k++ {
			fmt.Print("*")
		}

		// Second tree
		// for j := 1; j <= rows*2 - i*2; j++ {
		// 	fmt.Print(" ")
		// }

		// for k := 1; k <= 2*i - 1; k++ {
		// 	fmt.Print("*")
		// }
		fmt.Println()
	}

	// USING RANGES
	 for i := range 10 {
		i++
		fmt.Println(i)
	 }
}