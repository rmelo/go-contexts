package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func order(w http.ResponseWriter, r *http.Request) {

	log.Println("order: [started]")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	select {
	case <-ctx.Done():
		err := ctx.Err()

		log.Print(err)

		http.Error(w, err.Error(), http.StatusInternalServerError)

	case <-time.After(time.Second * 3):

		log.Println("order: processing payment [...]")

		if err := get(ctx, "http://localhost:4000/payment"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("order: payment [ok]")

		log.Println("order: confirmed [ok]")
		log.Println("order: [finished]")

		fmt.Fprint(w, "order placed!")
	}
}
