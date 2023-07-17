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
	"polemica_scrapper/database/config"
	"polemica_scrapper/handlers"
	"time"
)

var (
	appConfig config.Config
	dbConf    config.DbConfig
)

func init() {

	config.ParseConfig(&appConfig)
	config.ReadEnv(&appConfig)

	dbConf = appConfig.DbConfig

	m, err := migrate.New(
		"file://database/migrations/",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Dbname),
	)
	if err != nil {
		log.Fatalf("Cant handle migrations with err: %v", err)
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func main() {

	//init logger
	l := log.New(os.Stdout, "polemica-scr-back", log.LstdFlags)

	fmt.Println("init migration")

	//create new serve mux and register handlers
	sm := mux.NewRouter()

	//create handlers
	syncGameHandler := handlers.NewGameSync(l)
	syncStatsHandler := handlers.NewStatsHandler(l)

	//sync router
	syncRouter := sm.Methods(http.MethodGet).Subrouter()

	//register game endpoint
	syncRouter.HandleFunc("/register/{gameId}", syncGameHandler.RegisterGame).Methods(http.MethodGet)
	//fetch stats
	syncRouter.HandleFunc("/rating", syncStatsHandler.FetchRatings).Methods(http.MethodGet)

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
			dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.Dbname)

		db, err := sql.Open("postgres", connStr)
		l.Printf("Successful db connect: %v", db)

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
