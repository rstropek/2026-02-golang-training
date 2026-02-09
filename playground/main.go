package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func div(a, b int) int {
	return a / b
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter first number (or 'quit' to exit): ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)

		fmt.Print("Enter second number: ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)

		a, err1 := strconv.Atoi(input1)
		b, err2 := strconv.Atoi(input2)

		if err1 != nil || err2 != nil {
			fmt.Println("Error: Please enter valid integers")
			continue
		}

		result := div(a, b)
		fmt.Printf("%d / %d = %d\n\n", a, b, result)
	}
}
