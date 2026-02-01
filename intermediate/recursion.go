package basics

import "fmt"

func main() {
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(12345))

	fibonacci(10)
}

func factorial(n int) int {
	//  base case
	if n == 0 {
		return 1
	}
	// recursive case
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumOfDigits(n/10)
}

func fibonacci(n int) int {
	var a, b int = 0, 1

	for i := 0; i < n; i++ {
		fmt.Println(a, " ")
		a, b = b, a+b
	}
	return a
}
