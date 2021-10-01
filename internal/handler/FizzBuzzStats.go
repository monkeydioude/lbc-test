package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/monkeydioude/lbc-test/internal/entity"
	"github.com/monkeydioude/lbc-test/pkg/bbolt"
	"github.com/monkeydioude/lbc-test/pkg/response"
)

type stats struct {
	Params entity.Params `json:"parameters"`
	Hits   int           `json:"hits"`
}

// retrieveParamsAndHits iterate over entries recorded in boltDB's
// "stats" bucket and retrieves serialized parameters and hits
// of the most requested entry
func retrieveParamsAndHits() ([]byte, int, error) {
	var rawParams []byte
	hits := 0

	err := bbolt.Iterate("stats", func(key, value []byte) {
		h, err := strconv.Atoi(string(value))
		if err != nil {
			return
		}

		if h <= hits {
			return
		}

		hits = h
		rawParams = key
	})

	return rawParams, hits, err
}

// FizzBuzzStatsHandler GET /fizz-buzz/stats displays
// the parameters and the hit count of the most used GET /fizz-buzz query
func FizzBuzzStatsHandler(w http.ResponseWriter, req *http.Request) {
	// retrieve params and hits from DB
	rawParams, hits, err := retrieveParamsAndHits()

	if err != nil {
		log.Printf("[ERR ] could not retrieve stats: %s\n", err)
		response.InternalServer(w)
		return
	}

	params := entity.Params{}
	// unmarshaling result from DB. A marshaled entity.Params struct was used
	// as key in the DB
	err = json.Unmarshal(rawParams, &params)

	if err != nil {
		log.Printf("[ERR ] could not parse query: %s\n", err)
		response.InternalServer(w)
		return
	}

	// then building response by marshaling stats struct
	res, err := json.Marshal(stats{
		Params: params,
		Hits:   hits,
	})

	if err != nil {
		log.Printf("[ERR ] could not marshal response: %s\n", err)
		response.InternalServer(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response.Ok(w, res)
}
