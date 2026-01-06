package basics

import "fmt"

const pi = 3.14
const GRAVITY = 9.8

func main() {
	const days int = 7

	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
		sunday    = 7
	)

	fmt.Println("Value of pi: ", pi, "\nValue of Gravity: ", GRAVITY, "\nDays in a week: ", days,
		"\nDays of the week: ", monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}