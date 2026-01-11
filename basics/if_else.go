package basics

import "fmt"

func main() {
	temperature := 25

	if temperature > 30 {
		fmt.Println("It is hot outside")
	} else {
		fmt.Println("It is not hot outside")
	}

	grade := 85
	if grade >= 90 {
		fmt.Println("You got an A")
	} else if grade >= 80 {
		fmt.Println("You got a B")
	} else if grade >= 70 {
		fmt.Println("You got a C")
	} else if grade >= 60 {
		fmt.Println("You got a D")
	} else {
		fmt.Println("You got an F")
	}
}