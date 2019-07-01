package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
	defer cancel()

	log.Println("Testing context with deadline...")
	sayHello(ctx, "Hello", 2*time.Second)       // Will work
	sayHello(ctx, "World!", 5*time.Second)      // Will work
	sayHello(ctx, "Will fail!", 11*time.Second) // Will fail
}

func sayHello(ctx context.Context, msg string, d time.Duration) {
	select {
	case <-time.After(d):
		log.Println("Message: ", msg)
	case <-ctx.Done():
		log.Println("Error: ", ctx.Err())
	}
}
