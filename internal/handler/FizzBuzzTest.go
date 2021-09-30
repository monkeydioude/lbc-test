package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/monkeydioude/lbc-test/pkg/response"
)

type params struct {
	int1  int
	int2  int
	inter int
	limit int
	str1  string
	str2  string
}

// buildParamsFromValues construct a struct containing each mandatory values.
func buildParamsFromValues(values url.Values) (params, error) {
	p := params{
		str1: values.Get("str1"),
		str2: values.Get("str2"),
	}

	if p.str1 == "" || p.str2 == "" {
		return p, errors.New("str1 or str2 parameter missing")
	}

	if v, err := strconv.Atoi(values.Get("int1")); err == nil {
		p.int1 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("int2")); err == nil {
		p.int2 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("limit")); err == nil {
		p.limit = v
	} else {
		return p, err
	}

	return p, nil
}

// findOutToken is an algorithm generating the correct string token
// with respect to a given number and parameters provided to the endpoint
func findOutToken(p params, num int) string {
	if num%p.inter == 0 {
		return p.str1 + p.str2
	}

	if num%p.int1 == 0 {
		return p.str1
	}

	if num%p.int2 == 0 {
		return p.str2
	}

	return strconv.Itoa(num)
}

// FizzBuzzTestHandler is the handler for the route /fizz-buzz/test
func FizzBuzzTestHandler(w http.ResponseWriter, req *http.Request) {
	// building mandatory params struct from query string values
	params, err := buildParamsFromValues(req.URL.Query())
	if err != nil {
		response.BadRequest(w)
		return
	}

	// case of given number is equal int1 and int2
	params.inter = params.int1 * params.int2
	// setup response writer buffer
	resBuff := bytes.Buffer{}

	for i := 1; i <= params.limit; i++ {
		resBuff.WriteString(findOutToken(params, i) + ",")
	}

	// triming trailing "," before giving responseq
	response.Ok(w, bytes.Trim(resBuff.Bytes(), ","))
}
