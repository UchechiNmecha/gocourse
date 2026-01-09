package basics

import (
	"fmt"
	"math"
)

func main() {
	// Variable declaration and initialization

	var a, b int = 10, 3
	var result int

	// Addition
	result = a + b
	fmt.Println("Addition: ", result)

	// subtraction
	result = a - b
	fmt.Println("Subtraction: ", result)

	// Multipliication
	result = a * b
	fmt.Println("Multiplication: ", result)

	// Division
	result = a / b
	fmt.Println("Division: ", result)

	// Modulus
	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7.0
	fmt.Println("value of p: ", p)

	// Overflow with signed integers
	var maxInt68 int64 = 9223372036854775807
	fmt.Println("Max Int64: ", maxInt68)

	maxInt68 += 1
	fmt.Println("Maxint64 Overflow: ", maxInt68)


	// OverFlow with unsigned integers
	var maxUint64 uint64 = 18446744073709551615
	fmt.Println("Max Uint64: ", maxUint64)

	maxUint64 += 1
	fmt.Println("MaxUint64 Overflow: ", maxUint64)

	// Underflow
	var smallFloat float64 = 1.0e-323
	fmt.Println("small Float: ", smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}