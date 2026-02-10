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

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello(&wg)
	wg.Add(1)
	go sayHello(&wg)

	wg.Wait()
}
