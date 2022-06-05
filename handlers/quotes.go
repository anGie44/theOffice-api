package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anGie44/theOffice-api/data"
	"github.com/gorilla/mux"
)

type Quotes struct {
	l        *log.Logger
	quotesDB *data.QuotesDB
}

func NewQuotes(l *log.Logger, qdb *data.QuotesDB) *Quotes {
	return &Quotes{l, qdb}
}

type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func getEpisode(r *http.Request) (int, error) {
	vars := mux.Vars(r)

	episode, err := strconv.Atoi(vars["episode"])
	if err != nil {
		return 0, err
	}

	return episode, nil
}

func getSeason(r *http.Request) (int, error) {
	vars := mux.Vars(r)

	season, err := strconv.Atoi(vars["season"])
	if err != nil {
		return 0, err
	}

	return season, nil
}

func getFormat(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["format"]
}
