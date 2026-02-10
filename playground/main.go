package main

import (
	"fmt"
	"time"
)

func doSomethingAsync(result chan string) {
	defer close(result)
	fmt.Println("Starting to do something")
	result <- "Step 1"
	fmt.Println("Step 1 done")
	time.Sleep(1 * time.Second)
	result <- "Step 2"
	fmt.Println("Step 2 done")
	fmt.Println("We are done")
}

func main() {
	result := make(chan string)
	go doSomethingAsync(result)
	time.Sleep(10 * time.Second)
	fmt.Println(<-result)
	fmt.Println(<-result)
}
