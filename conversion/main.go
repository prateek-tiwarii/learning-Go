package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome!!!")
	fmt.Println("Enter the rating for the restaurant:")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	if err != nil {
		fmt.Println(err)
	}

	increment, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	fmt.Println("The incremented ratings is:", increment+1)

}
