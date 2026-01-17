package basics

import "fmt"

func main() {
	process(10)

}

func process(i int) {
	defer fmt.Println("Deferred value of i: ", i)
	defer fmt.Println("First Deferred statement executed")
	defer fmt.Println("Second Deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i before function ends: ", i)
}
