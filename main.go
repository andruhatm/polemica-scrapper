package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	//init logger
	l := log.New(os.Stdout, "polemica-scr-back", log.LstdFlags)

	fmt.Println("hello world")

	//create new serve mux and register handlers
	sm := mux.NewRouter()

	//sync router
	syncRouter := sm.Methods(http.MethodGet).Subrouter()
	//register game endpoint
	syncRouter.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write([]byte("OK"))
		if err != nil {
			l.Printf("Error while writing the data to an HTTP reply with err=%s", err)
			return
		}
	})

	//probes router
	probesRouter := sm.Methods(http.MethodGet).Subrouter()
	probesRouter.HandleFunc("/probes/readiness", func(rw http.ResponseWriter, r *http.Request) {
		_, err := rw.Write([]byte("OK"))
		if err != nil {
			l.Printf("Error while writing the data to an HTTP reply with err=%s", err)
			return
		}
	})
	probesRouter.HandleFunc("/probes/liveness", func(rw http.ResponseWriter, r *http.Request) {
		//TODO check if we can access DB
		_, err := rw.Write([]byte("OK"))
		if err != nil {
			l.Printf("Error while writing the data to an HTTP reply with err=%s", err)
			return
		}
	})

	s := http.Server{
		Addr:         ":8081",
		Handler:      sm,
		TLSConfig:    nil,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 8081")

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//trap os.Signal and gracefully shutdown the server
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	l.Printf("Graceful shutdown with signal %s \n", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)

}
