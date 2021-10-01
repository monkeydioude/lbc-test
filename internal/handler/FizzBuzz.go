package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/monkeydioude/lbc-test/internal/entity"
	"github.com/monkeydioude/lbc-test/pkg/bbolt"
	"github.com/monkeydioude/lbc-test/pkg/response"
)

// buildParamsFromValues construct a struct containing each mandatory values.
func buildParamsFromValues(values url.Values) (entity.Params, error) {
	p := entity.Params{
		Str1: values.Get("str1"),
		Str2: values.Get("str2"),
	}

	if p.Str1 == "" || p.Str2 == "" {
		return p, errors.New("str1 or str2 parameter missing")
	}

	if v, err := strconv.Atoi(values.Get("int1")); err == nil {
		p.Int1 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("int2")); err == nil {
		p.Int2 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("limit")); err == nil {
		p.Limit = v
	} else {
		return p, err
	}

	// number being a multiple of Int1 and Int2
	if p.Int1 == p.Int2 {
		p.Inter = p.Int1
	} else {
		p.Inter = p.Int1 * p.Int2
	}
	return p, nil
}

// findOutToken is an algorithm generating the correct string token
// with respect to a given number and parameters provided to the endpoint
func findOutToken(p entity.Params, num int) string {
	if num%p.Inter == 0 {
		return p.Str1 + p.Str2
	}

	if num%p.Int1 == 0 {
		return p.Str1
	}

	if num%p.Int2 == 0 {
		return p.Str2
	}

	return strconv.Itoa(num)
}

// computeResponse build a response by running the main algorithm
// through every int from 1 to Limit
func computeResponse(p entity.Params) []byte {
	// setup response writer buffer
	resBuff := bytes.Buffer{}

	for i := 1; i <= p.Limit; i++ {
		resBuff.WriteString(findOutToken(p, i) + ",")
	}

	// triming trailing "," before giving response back
	return bytes.Trim(resBuff.Bytes(), ",")
}

// recordHit increments a counter defining the number of times
// a POST /fizz-buzz request, and its specific parameters, has been made.
// We use the marshaled version of the entity.Params struct as key of
// "stats" bucket
func recordHit(query string) error {
	record, err := bbolt.Fetch("stats", query)

	if err != nil || len(record) == 0 {
		log.Printf("[WARN] could not fetch %s\n", query)
		record = []byte("0")
	}

	r, err := strconv.Atoi(string(record))
	if err != nil {
		return err
	}
	r += 1

	err = bbolt.Write("stats", query, strconv.Itoa(r))

	if err != nil {
		return err
	}

	return nil
}

// FizzBuzzTestHandler is the handler for the route /fizz-buzz/test
func FizzBuzzTestHandler(w http.ResponseWriter, req *http.Request) {
	// building mandatory entity.Params struct from query string values
	params, err := buildParamsFromValues(req.URL.Query())

	if err != nil {
		log.Printf("[ERR ] could not build parametrers: %s\n", err)
		response.BadRequest(w)
		return
	}

	pm, err := json.Marshal(params)
	if err == nil {
		if err = recordHit(string(pm)); err != nil {
			log.Printf("[ERR ] could not record hit into DB: %s\n", err)
		}
	}

	response.Ok(w, computeResponse(params))
}
