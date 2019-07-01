package main

import (
	"context"
	"log"
)

type color string

func main() {
	ctx := context.Background()

	black := color("black")

	// Use context Values only for request-scoped data that transits
	// processes and APIs, not for passing optional parameters to functions.
	ctx = context.WithValue(ctx, black, "color")

	log.Println("Testing context with value...")
	showColor(ctx, black)
	showColor(ctx, "red")
}

func showColor(ctx context.Context, c color) {
	if value := ctx.Value(c); value != nil {
		log.Println("Color found: ", c)
		return
	}
	log.Println("Color not found: ", c)
}
