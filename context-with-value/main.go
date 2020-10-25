package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "userID", "user.2020")
	if err := handle(ctx); err != nil {
		panic(err)
	}
}

func handle(ctx context.Context) error {
	 userID := ctx.Value("userID")
	 fmt.Printf("Executing handler with userID %s\n", userID)
	return update(ctx)
}

func update(ctx context.Context) error {
	userID := ctx.Value("userID")
	fmt.Printf("Update user with userID %s\n", userID)
	return nil
}