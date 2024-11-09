package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("reading input from the user")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings for the restaurant:")
	input, _ := reader.ReadString('\n')

	fmt.Println("hey, thanks for rating", input)

}
