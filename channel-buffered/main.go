package main

import (
	"fmt"
	"time"
)

func doSomethingAsync(result chan string) {
	defer close(result)
	fmt.Println("Starting to do something")
	result <- "Step 1" // No longer a blocking operation because the channel is buffered
	fmt.Println("Step 1 done")
	time.Sleep(1 * time.Second)
	result <- "Step 2"
	fmt.Println("Step 2 done")
	fmt.Println("We are done")
}

func main() {
	result := make(chan string, 2)
	go doSomethingAsync(result)
	time.Sleep(10 * time.Second)
	fmt.Println(<-result) // Blocking operation (wait for sender to send a value)
	fmt.Println(<-result)
}
