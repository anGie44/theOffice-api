package models

import (
	combinations "github.com/mxschmitt/golang-combinations"
)

// swagger:model
type Connection struct {
	Episode     int     `bson:"episode" json:"episode"`
	EpisodeName string  `bson:"episode_name" json:"episode_name"`
	Links       []*Link `bson:"links" json:"links"`
	Nodes       []*Node `bson:"nodes" json:"nodes"`
}

// swagger:model
type Link struct {
	Source string `bson:"source" json:"source"`
	Target string `bson:"target" json:"target"`
	Value  int    `bson:"value" json:"value"`
}

// swagger:model
type Node struct {
	Id string `bson:"id" json:"id"`
}

type Connections []*Connection

type episode struct {
	name   string
	scenes map[int]Quotes
}

func GetConnectionsPerEpisode(seasonQuotes Quotes) Connections {
	// Map of quotes per episode
	quotesByEpisode := make(map[int]*episode)

	for _, quote := range seasonQuotes {
		_, ok := quotesByEpisode[quote.Episode]
		if ok {
			_, sceneOk := quotesByEpisode[quote.Episode].scenes[quote.Scene]
			if sceneOk {
				quotesByEpisode[quote.Episode].scenes[quote.Scene] = append(quotesByEpisode[quote.Episode].scenes[quote.Scene], quote)
			} else {
				quotesByEpisode[quote.Episode].scenes[quote.Scene] = Quotes{quote}
			}
		} else {
			quotesByEpisode[quote.Episode] = &episode{
				name: quote.EpisodeName,
				scenes: map[int]Quotes{
					quote.Scene: []*Quote{quote},
				},
			}
		}
	}

	var connections Connections

	for episodeNum, episode := range quotesByEpisode {
		// Loop through scenes in an episode
		linksMap := make(map[Link]int)
		nodesMap := make(map[string]struct{})

		for _, quotes := range episode.scenes {
			charactersInScene := make(map[string]struct{})
			var exists = struct{}{}

			for _, quote := range quotes {
				charactersInScene[quote.Character] = exists
			}

			var chars []string
			for c := range charactersInScene {
				chars = append(chars, c)
				nodesMap[c] = exists
			}

			if len(chars) < 2 {
				continue
			}

			combinations := combinations.Combinations(chars, 2)

			for _, combo := range combinations {
				link := Link{
					Source: combo[0],
					Target: combo[1],
				}

				linksMap[link]++
			}
		}

		var nodes []*Node
		for c := range nodesMap {
			nodes = append(nodes, &Node{
				Id: c,
			})
		}

		var links []*Link
		for l, count := range linksMap {
			links = append(links, &Link{
				Source: l.Source,
				Target: l.Target,
				Value:  count,
			})
		}

		connections = append(connections, &Connection{
			Episode:     episodeNum,
			EpisodeName: episode.name,
			Links:       links,
			Nodes:       nodes,
		})
	}

	return connections
}
