package main

import "fmt"

func isEvenLength(s string) bool {
	characterCount := len(s)
	if characterCount%2 == 0 {
		fmt.Printf("%d - has an even number of characters\n", characterCount)
		return true
	} else {
		fmt.Printf("%d - has an odd number of characters\n", characterCount)
		return false
	}

}

func main() {
	fmt.Println("Hello, world.")
}
