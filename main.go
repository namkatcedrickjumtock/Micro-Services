package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/namkatcedrickjumtock/Micro-services/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api:- ", log.LstdFlags)
	hl := handlers.NewHello(l)

	// creatings a new serve Mux which also implements the HandleFucn interface
	sm := http.NewServeMux()

	// register a handler
	sm.Handle("/", hl)

	// http.ListenAndServe(":3000", sm)

	// creating an http server with secure additional params
	s := &http.Server{
		Addr:         ":3000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// go Routine func
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Received, Terminate Gracefule shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	
	s.Shutdown(tc)
}
