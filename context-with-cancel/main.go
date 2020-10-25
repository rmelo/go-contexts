package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f := func(n int, ctx context.Context) {
		fmt.Printf("func %v started with %v sec\n", n, n)
		select {
		case <-time.After(time.Second * time.Duration(n)):
			fmt.Printf("func %v win!\t\t[ok]\n", n)
			cancel()
		case <-ctx.Done():
			fmt.Printf("func %v\t\t[cancelled]\n", n)
			return
		}
	}

	go f(1, ctx)
	go f(2, ctx)

	fmt.Println("Finishing program...\t[WAITING]")
	time.Sleep(time.Second * time.Duration(5))
	fmt.Println("Finished\t\t[OK]")
}
