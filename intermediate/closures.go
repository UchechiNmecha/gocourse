package basics

import "fmt"

func main() {

	sequence := adder()

	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	sequence2 := adder()
	fmt.Println(sequence2())

	subtracter := func() func(int) int {
		sum := 99
		return func(x int) int {
			sum -= x
			return sum
		}
	}()

	fmt.Println(subtracter(10))
	fmt.Println(subtracter(20))
	fmt.Println(subtracter(5))
}

func adder() func() int {
	i := 0
	fmt.Println("This is the initial value of i: ", i)

	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
