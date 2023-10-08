package handlers

import (
	"net/http"

	"github.com/anGie44/theOffice-api/helpers"
	"github.com/anGie44/theOffice-api/models"
)

// swagger:route GET /quotes quotes getQuotes
// Return a list of quotes from the database
// responses:
//		200: quotesResponse

// GetQuotes handles GET requests and returns all quotes
func (q *Quotes) GetQuotes(rw http.ResponseWriter, r *http.Request) {
	q.l.Println("Handle GET Quotes")
	rw.Header().Add("Content-Type", "application/json")

	quotes, err := q.quotesDB.GetQuotes()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		helpers.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = helpers.ToJSON(quotes, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /season/{season}/format/{format} interface getQuotesBySeasonWithFormat
// Return a list of quotes from the database for a given season and format
// responses:
//	 	200: formattedQuotesResponse

// GetQuotesBySeasonWithFormat handles GET requests and returns quotes for the specified season and format
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
	rw.Header().Add("Content-Type", "application/json")

	quotes, err := q.quotesDB.GetQuotesBySeason(season)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		helpers.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if format == "quotes" {
		err = helpers.ToJSON(quotes, rw)
		if err != nil {
			http.Error(rw, "Unable to marshal quotes json", http.StatusInternalServerError)
		}
		return
	}

	// Connections

	connections := models.GetConnectionsPerEpisode(quotes)

	err = helpers.ToJSON(connections, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal connections json", http.StatusInternalServerError)
	}
}

// swagger:route GET /season/{season}/episode/{episode} quotes getQuotesBySeasonAndEpisode
// Return a list of quotes from the database for a given season and episode
// responses:
//		200: quotesResponse

// GetQuotesBySeasonAndEpisode handles GET requests and returns quotes for the specified season and episode
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
	rw.Header().Add("Content-Type", "application/json")

	quotes, err := q.quotesDB.GetQuotesBySeasonAndEpisode(season, episode)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		helpers.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = helpers.ToJSON(quotes, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
