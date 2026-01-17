package basics

import "fmt"

func main() {
	statement, total := sum("The total of (1, 2, 3, 4, 5) is: ", 1, 2, 3, 4, 5)
	fmt.Println(statement, total)

	numbers := []int{10, 20, 30}
	statement, total = sum("The total of (10, 20, 30) is: ", numbers...)
	fmt.Println(statement, total)
}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}
