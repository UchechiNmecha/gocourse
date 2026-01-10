package basics

func main() {

	i := 1
	for i <= 5 {
		println("Number:", i)
		i++
	}

	// Infinite loop example:
	// for {
	// 	println("Infinite Loop")
	// }

	// Simulating a while loop with an infinite loop and break
	sum := 0
	for {
		sum += 10
		println("Sum:", sum)
		if sum == 50 {
			break
		}
	}

	// For as while with continue
	num := 0

	for {
		num++
		if num%2 == 0 {
			continue
		}

		if num >= 10 {
			break
		}
		println("Odd Number: ", num)
	}
}