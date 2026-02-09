package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter first number: ")
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

		result, err := div(a, b)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("%d / %d = %d\n\n", a, b, result)
	}
}
