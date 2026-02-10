package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func div(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func repl() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("a: ")
		a, _ := reader.ReadString('\n')
		fmt.Print("b: ")
		b, _ := reader.ReadString('\n')
		a = strings.TrimSpace(a)
		b = strings.TrimSpace(b)
		aInt, _ := strconv.Atoi(a)
		bInt, _ := strconv.Atoi(b)
		result := div(aInt, bInt)
		fmt.Printf("result: %d\n", result)
	}
}

func main() {
	for {
		repl()
	}
}
