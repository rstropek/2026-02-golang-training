package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello, World!")

	// Simulate some work by sleeping for a bit
	time.Sleep(500 * time.Millisecond)
}

func getValueAsync(result chan int) {
	defer close(result)

	result <- 42
	time.Sleep(500 * time.Millisecond)
	//result <- 43
	result <- 44
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()

	result := make(chan int)
	go getValueAsync(result)

	// Consume exactly one value
	//value := <-result
	//fmt.Println("Value:", value)
	
	// Consume a variable number of values
	// NOTE: Do not forget to CLOSE the channel
	//for value := range result {
	//	fmt.Println("Value:", value)
	//}

	// Conditionally consume values
	time.Sleep(25 * time.Millisecond)
	select {
	case value := <-result: // No longer a blocking operation
		fmt.Println("Value:", value)
	default:
		fmt.Println("No value available")
	}
}
