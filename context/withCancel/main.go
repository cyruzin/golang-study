package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Println("Testing context with cancel...")
	sayHello(ctx, "Hello", 2*time.Second)  // Will work
	sayHello(ctx, "World!", 5*time.Second) // Will work
}

func sayHello(ctx context.Context, msg string, d time.Duration) {
	select {
	case <-time.After(d):
		log.Println("Message: ", msg)
	case <-ctx.Done():
		return
	}
}
