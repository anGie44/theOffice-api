package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrQuotesNotFound = fmt.Errorf("Quotes not found")

// Quote defines the structure for an API quote
// swagger:model
type Quote struct {
	// the id for the quote
	Id int `bson:"-" json:"-"`

	// the season the quote is from
	Season int `bson:"season" json:"season"`

	// the episode the quote is from
	Episode int `bson:"episode" json:"episode"`

	// the scene the quote is from
	Scene int `bson:"scene" json:"scene"`

	// the episode name the quote is from
	EpisodeName string `bson:"episode_name" json:"episode_name"`

	// the character the quote is associated with
	Character string `bson:"character" json:"character"`

	// the quote
	Quote string `bson:"quote" json:"quote"`
}

type Quotes []*Quote

type QuotesDB struct {
	ctx  context.Context
	opts *QuotesDBOptions
	l    *log.Logger
}

type QuotesDBOptions struct {
	hostname   string
	database   string
	username   string
	password   string
	collection string
}

func NewQuotesDBOptions(hostname, database, username, password, collection string) *QuotesDBOptions {
	return &QuotesDBOptions{hostname, database, username, password, collection}
}

func NewQuotesDB(opts *QuotesDBOptions, l *log.Logger) *QuotesDB {
	return &QuotesDB{opts: opts, l: l}
}

// GetQuotes return a random set of quotes from the database
func (q *QuotesDB) GetQuotes() (Quotes, error) {

	return nil, nil
}

// GetQuotesBySeason returns all quotes for the given season from the database
func (q *QuotesDB) GetQuotesBySeason(season int) (Quotes, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", q.opts.username, q.opts.password, q.opts.hostname)).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	q.ctx = ctx

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)

	db := client.Database(q.opts.database)
	collection := db.Collection(q.opts.collection)

	filterCursor, err := collection.Find(q.ctx, bson.M{"season": season})
	if err != nil {
		q.l.Printf("error while searching for quotes in db: %s", err)
		return nil, err
	}

	var results Quotes
	for filterCursor.Next(q.ctx) {
		var quote Quote
		err := filterCursor.Decode(&quote)
		if err != nil {
			q.l.Printf("error while decoding filtered quote: %s", err)
			continue
		}

		results = append(results, &quote)

	}

	return results, nil
}

// GetQuotesBySeasonAndEpisode returns all quotes for the given season and episode from the database
func (q *QuotesDB) GetQuotesBySeasonAndEpisode(season, episode int) (Quotes, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s", q.opts.username, q.opts.password, q.opts.hostname)).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	q.ctx = ctx

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	defer client.Disconnect(ctx)

	db := client.Database(q.opts.database)
	collection := db.Collection(q.opts.collection)

	filterCursor, err := collection.Find(q.ctx, bson.M{"season": season, "episode": episode})
	if err != nil {
		q.l.Printf("error while searching for quotes in db: %s", err)
		return nil, err
	}

	var results Quotes
	for filterCursor.Next(q.ctx) {
		var quote Quote
		err := filterCursor.Decode(&quote)
		if err != nil {
			q.l.Printf("error while decoding filtered quote: %s", err)
			continue
		}

		results = append(results, &quote)

	}

	return results, nil
}
