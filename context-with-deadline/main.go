package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	deadline := time.Now().Add(time.Duration(time.Second * 2))

	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	select {
		case <-time.After(2 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
	}

}
