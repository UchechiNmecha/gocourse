package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// The rest of the guessing game logic would go here

	// welcome Message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Can you guess it?")
	fmt.Println("The target number is:", target) // For testing purposes

	var guess int

	for {
		fmt.Println("Enter your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number")
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}