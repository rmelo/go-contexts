package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f := func(n int, ctx context.Context) {
		rand.Seed(time.Now().UnixNano())
		fmt.Printf("func %v started with %v sec\n", n, n)
		select {
		case <-time.After(time.Second * time.Duration(rand.Intn(5))):
			fmt.Printf("func %v win!\t\t[ok]\n", n)
			cancel()
		case <-ctx.Done():
			fmt.Printf("func %v\t\t[cancelled]\n", n)
			return
		}
	}

	go f(1, ctx)
	go f(2, ctx)
	go f(3, ctx)

	fmt.Println("Finishing program...\t[WAITING]")
	time.Sleep(time.Second * time.Duration(3))
	fmt.Println("Finished\t\t[OK]")
}
