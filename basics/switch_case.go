package basics

import "fmt"

func main() {
	// Basic switch case
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("This is an apple fruit")
	case "banana":
		fmt.Println("This a a banana")
	default:
		fmt.Println("Unknown fruit")
	}

	// Multiple cases in a single case statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is weekday")
	case "Saturday", "Sunday":
		fmt.Println("It is weekend")
	default:
		fmt.Println("unknown day")
	}

	// Switch with conditions
	number := 15
	switch {
	case number < 10:
		fmt.Println("number is less than 10")
	case number >= 10 && number <= 20:
		fmt.Println("number is between 10 and 20")
	default:
		fmt.Println("number is greater than 20")
	}

	// Using fallthrough
	num := 2
	switch {
	case num > 1:
		fmt.Println("num is greater than 1")
		fallthrough
	case num == 2:
		fmt.Println("num is two")
	default:
		fmt.Println("num is something else")
	}

	// Type switch
	checkType(23)
	checkType("hello")
	checkType(3.14)
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("This is an Integer")
	case string:
		fmt.Println("This is a string")
	case float64:
		fmt.Println("This is a float64")
	default:
		fmt.Println("Unknown type")
	}
}
