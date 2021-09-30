package handler

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/monkeydioude/lbc-test/pkg/response"
)

type params struct {
	int1  int
	int2  int
	limit int
	str1  string
	str2  string
}

// buildParamsFromValues construct a struct containing each mandatory values
func buildParamsFromValues(values url.Values) (params, error) {
	p := params{
		str1: values.Get("str1"),
		str2: values.Get("str2"),
	}

	if p.str1 == "" || p.str2 == "" {
		return p, errors.New("str1 or str2 parameter missing")
	}

	if v, err := strconv.Atoi(values.Get("int1")); err != nil {
		p.int1 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("int2")); err != nil {
		p.int2 = v
	} else {
		return p, err
	}

	if v, err := strconv.Atoi(values.Get("limit")); err != nil {
		p.limit = v
	} else {
		return p, err
	}

	return p, nil
}

func FizzBuzzTestHandler(w http.ResponseWriter, req *http.Request) {
	params, err := buildParamsFromValues(req.URL.Query())
	if err != nil {
		response.BadRequest(w)
		return
	}

	intersect := params.int1 * params.int2
}
