package basics

import "fmt"

func main() {
	recoverFromPanic()
	fmt.Println("Returned From process")
}

func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic, e.g., log it
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Code that may cause a panic
	fmt.Println("About to panic")
	panic("Something went wrong!")
	fmt.Println("End process")
}
