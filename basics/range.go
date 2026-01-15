package basics

import "fmt"

func main() {
	var message string = "Hello"
	for index, char := range message {
		fmt.Println("Index:", index, "Character:", string(char))
		fmt.Printf("Index: %d, Rune: %c\n", index, char)
	}
}
