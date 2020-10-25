package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	timeout := time.Duration(time.Second * 3)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	select {
		case <- time.After(2 * time.Second):
			fmt.Println("overslept")
		case <- ctx.Done():
			fmt.Println(ctx.Err())
	}

}