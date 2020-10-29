package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func antifraud(w http.ResponseWriter, r *http.Request) {

	log.Println("antifraud: [started]")
	ctx := r.Context()

	select {

	case <-ctx.Done():

		err := ctx.Err()
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	case <-time.After(time.Second * 3):

		log.Println("antifraud: [approved]")
		log.Println("antifraud: [finished]")
		fmt.Fprint(w, "approved!")
	}
}
