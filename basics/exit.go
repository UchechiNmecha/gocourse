package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("This will not be printed because os.Exit is called before it")

	fmt.Println("Starting the main function")

	// Exiting the program
	os.Exit(1)
	// This line will not be executed
	fmt.Println("End of the main function")
}
