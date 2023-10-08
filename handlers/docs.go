package handlers

import "github.com/anGie44/theOffice-api/models"

// A list of quotes
// swagger:response quotesResponse
type quotesResponseWrapper struct {
	Body []models.Quote
}

// A list of quotes or connections
// swagger:response formattedQuotesResponse
type formattedQuotesResponseWrapper struct {
	Body interface{}
}
