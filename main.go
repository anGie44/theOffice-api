package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/anGie44/theOffice-api/data"
	"github.com/anGie44/theOffice-api/handlers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var dbHost = env.String("MONGODB_HOST", false, "localhost", "mongodb host address")
var dbName = env.String("MONGODB_DATABASE", false, "the-office", "mongodb database name")
var dbUsername = env.String("MONGODB_USERNAME", false, "boss", "mongodb database username")
var dbPassword = env.String("MONGODB_PASSWORD", false, "password", "mongodb database user password")
var dbCollection = env.String("MONGODB_COLLECTION", false, "quotes", "mongodb database collection")

func main() {
	var bindAddress string
	port := os.Getenv("PORT") // value provided by Heroku
	if port == "" {
		bindAddress = ":8080"
	} else {
		bindAddress = fmt.Sprintf(":%s", port)
	}

	env.Parse()

	l := log.New(os.Stdout, "theOffice-api ", log.LstdFlags)

	// Create MongoDB Client
	dbOpts := data.NewQuotesDBOptions(*dbHost, *dbName, *dbUsername, *dbPassword, *dbCollection)
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
		Addr:        bindAddress,
		Handler:     ch(sm),
		ErrorLog:    l,
		ReadTimeout: 1 * time.Minute,
		IdleTimeout: 2 * time.Minute,
	}

	go func() {
		l.Printf("Starting server on port %s\n", strings.TrimPrefix(bindAddress, ":"))

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
