package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	port := flag.String("p", "3000", "server port")
	flag.Parse()

	srv := &http.Server{
		Addr: fmt.Sprintf(":%v", *port),
	}

	fmt.Printf("\n\nMagic happens on http://localhost:%v\n\n", *port)

	http.HandleFunc("/order", order)
	http.HandleFunc("/payment", payment)
	http.HandleFunc("/antifraud", antifraud)

	go func() {
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
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
