package main

import (
	"context"
	"fmt"
	"time"
)

func sayHello(ctx context.Context) {
	select {
		case <-time.After(5 * time.Second):
			fmt.Println("Hello, World!")
		case <-ctx.Done():
			fmt.Println("Context cancelled, timeout")
			return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()

	sayHello(ctx)
}
