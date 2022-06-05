package handlers

import (
	"net/http"

	"github.com/anGie44/theOffice-api/data"
)

// GetQuotes handles GET requests and returns a random subset of quotes
func (q *Quotes) GetQuotes(rw http.ResponseWriter, r *http.Request) {
	q.l.Println("Handle GET Quotes")

	quotes, err := q.quotesDB.GetQuotes()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(quotes, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// GetQuotesBySeason handles GET requests and returns quotes for the specified season and format
// GET /season/{season}/format/{format}
func (q *Quotes) GetQuotesBySeasonWithFormat(rw http.ResponseWriter, r *http.Request) {
	season, err := getSeason(r)
	if err != nil {
		http.Error(rw, "Unable to convert season", http.StatusBadRequest)
		return
	}

	format := getFormat(r)
	if format == "" {
		http.Error(rw, "Must specify a format", http.StatusBadRequest)
		return
	}

	q.l.Printf("Handle GET Quotes for Season (%d) in Format (%s)\n", season, format)

	quotes, err := q.quotesDB.GetQuotesBySeason(season)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if format == "quotes" {
		err = data.ToJSON(quotes, rw)
		if err != nil {
			http.Error(rw, "Unable to marshal quotes json", http.StatusInternalServerError)
		}
		return
	}

	// Connections

	connections := data.GetConnectionsPerEpisode(quotes)

	err = data.ToJSON(connections, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal connections json", http.StatusInternalServerError)
	}
}

// GetQuotesBySeason handles GET requests and returns quotes for the specified season and episode
// GET /season/{season}/episode/{episode}
func (q *Quotes) GetQuotesBySeasonAndEpisode(rw http.ResponseWriter, r *http.Request) {
	season, err := getSeason(r)
	if err != nil {
		http.Error(rw, "Unable to convert season", http.StatusBadRequest)
		return
	}

	episode, err := getEpisode(r)
	if err != nil {
		http.Error(rw, "Unable to convert episode", http.StatusBadRequest)
		return
	}

	q.l.Printf("Handle GET Quotes for Season (%d) Episode (%d)\n", season, episode)

	quotes, err := q.quotesDB.GetQuotesBySeasonAndEpisode(season, episode)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(quotes, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
