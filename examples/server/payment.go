package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func payment(w http.ResponseWriter, r *http.Request) {

	log.Println("payment: [started]")

	ctx := r.Context()

	select {

	case <-ctx.Done():

		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	case <-time.After(time.Second * 3):

		log.Println("payment: checking antifraud system [...]")

		if err := get(ctx, "http://localhost:5000/antifraud"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		log.Println("payment: antifraud [ok]")

		log.Println("payment: [approved]")
		log.Println("payment: [finished]")
		fmt.Fprint(w, "payed!")
	}
}
