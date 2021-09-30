package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestICanReturnACorrectResponseOn2PositiveIntAndPositiveLimit(t *testing.T) {
	int1 := 2
	int2 := 5
	limit := 10

	str1 := "clair-de-"
	str2 := "lune"

	goal := "1,clair-de-,3,clair-de-,lune,clair-de-,7,clair-de-,9,clair-de-lune"

	target := fmt.Sprintf(
		"http://localhost:8004/fizz-buzz/test?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s",
		int1,
		int2,
		limit,
		str1,
		str2,
	)

	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", target, nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Body.String() != goal {
		t.Fail()
	}
}

func TestICanFailOnMissingParameters(t *testing.T) {
	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", "http://localhost:8004/fizz-buzz/test?int1=1", nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Result().StatusCode != 403 {
		t.Fail()
	}
}

func TestICanFailOnUnwantedIntValues(t *testing.T) {
	wtest := httptest.NewRecorder()
	// culprit here is limit=a
	reqtest := httptest.NewRequest("GET", "http://localhost:8004/fizz-buzz/test?int1=1&int2=2&limit=a&str1=a&str2=b", nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Result().StatusCode != 403 {
		t.Fail()
	}
}

func TestICanReturnNothingOnLimit0(t *testing.T) {
	int1 := 2
	int2 := 5
	limit := 0

	str1 := "clair-de-"
	str2 := "vide"

	goal := ""

	target := fmt.Sprintf(
		"http://localhost:8004/fizz-buzz/test?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s",
		int1,
		int2,
		limit,
		str1,
		str2,
	)

	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", target, nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Body.String() != goal {
		t.Fail()
	}
}

func TestICanReturnNothingOnNegativeLimit(t *testing.T) {
	int1 := 2
	int2 := 5
	limit := -1

	str1 := "clair-de-"
	str2 := "trou-noir"

	goal := ""

	target := fmt.Sprintf(
		"http://localhost:8004/fizz-buzz/test?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s",
		int1,
		int2,
		limit,
		str1,
		str2,
	)

	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", target, nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Body.String() != goal {
		t.Fail()
	}
}

func TestInt2LargerThanLimit(t *testing.T) {
	int1 := 2
	int2 := 5
	limit := 4

	str1 := "int2?"
	str2 := "*zZz*"

	goal := "1,int2?,3,int2?"

	target := fmt.Sprintf(
		"http://localhost:8004/fizz-buzz/test?int1=%d&int2=%d&limit=%d&str1=%s&str2=%s",
		int1,
		int2,
		limit,
		str1,
		str2,
	)

	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", target, nil)
	FizzBuzzTestHandler(wtest, reqtest)

	if wtest.Body.String() != goal {
		t.Fail()
	}
}
