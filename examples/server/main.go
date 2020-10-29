package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
)

func main() {

	port := flag.String("p", "3000", "server port")
	flag.Parse()

	fmt.Printf("\n\nMagic happens on http://localhost:%v\n\n", *port)

	http.HandleFunc("/order", order)
	http.HandleFunc("/payment", payment)
	http.HandleFunc("/antifraud", antifraud)

	http.ListenAndServe(fmt.Sprintf(":%v", *port), nil)
}

func get(ctx context.Context, endpoint string) error {

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
