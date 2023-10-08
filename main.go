package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anGie44/theOffice-api/data"
	"github.com/anGie44/theOffice-api/handlers"

	"github.com/caarlos0/env/v9"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type config struct {
	Port         int    `env:"PORT" envDefault:"8080"`
	DBHost       string `env:"MONGODB_HOST" envDefault:"localhost"`
	DBDatabase   string `env:"MONGODB_DATABASE" envDefault:"the-office"`
	DBCollection string `env:"MONGODB_COLLECTION" envDefault:"quotes"`
	DBUsername   string `env:"MONGODB_USERNAME" envDefault:"boss"`
	DBPassword   string `env:"MONGODB_PASSWORD" envDefault:"password"`
}

func main() {
	l := log.New(os.Stdout, "theOffice-api ", log.LstdFlags)

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		l.Printf("Error parsing environment config: %s\n", err)
		os.Exit(1)
	}

	// Create MongoDB Client
	dbOpts := data.NewQuotesDBOptions(cfg.DBHost, cfg.DBDatabase, cfg.DBUsername, cfg.DBPassword, cfg.DBCollection)
	db := data.NewQuotesDB(dbOpts, l)

	wh := handlers.NewWelcome(l)
	qh := handlers.NewQuotes(l, db)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", wh.Welcome)
	// Deprecated handlers
	getRouter.HandleFunc("/season/{season:[1-9]}/format/{format:quotes|connections}", qh.GetQuotesBySeasonWithFormat)
	getRouter.HandleFunc("/season/{season:[1-9]}/episode/{episode:[1-9]|[1][0-9]|2[0-3]}", qh.GetQuotesBySeasonAndEpisode)
	// V2 handlers to use request body for filtering data
	getRouter.HandleFunc("/v2/quotes", qh.GetQuotes)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	s := http.Server{
		Addr:        fmt.Sprintf(":%d", cfg.Port),
		Handler:     ch(sm),
		ErrorLog:    l,
		ReadTimeout: 1 * time.Minute,
		IdleTimeout: 2 * time.Minute,
	}

	go func() {
		l.Printf("Starting server on port %d\n", cfg.Port)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received
	sig := <-c
	log.Println("Received signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(ctx)
}
