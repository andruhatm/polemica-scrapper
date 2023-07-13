package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"polemica_scrapper/handlers"
	"time"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "games"
)

func init() {
	//db, _ := sql.Open("postgres", "postgres://db:5433/database?sslmode=enable")
	//driver, _ := postgres.WithInstance(db, &postgres.Config{})
	//m, _ := migrate.NewWithDatabaseInstance(
	//	"file:///migrations",
	//	"postgres", driver)
	//m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run

	m, err := migrate.New(
		"file://database/migrations/",
		"postgres://postgres:123@db:5433/games?sslmode=disable")
	// postgres://user:secret@localhost:5432/mydatabasename
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	//init logger
	l := log.New(os.Stdout, "polemica-scr-back", log.LstdFlags)

	fmt.Println("init migration")

	//create new serve mux and register handlers
	sm := mux.NewRouter()

	//create handler
	syncGameHandler := handlers.NewGameSync(l)

	//sync router
	syncRouter := sm.Methods(http.MethodGet).Subrouter()
	//register game endpoint
	syncRouter.HandleFunc("/register/{gameId}", syncGameHandler.RegisterGame).Methods(http.MethodGet)

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
		//check if we can access DB

		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		_, err := sql.Open("postgres", connStr)

		if err != nil {
			l.Printf("Error while connection to DB with err=%s", err)
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
