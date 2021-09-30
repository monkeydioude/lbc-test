package handler

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestICanReturnsACorrectResponseOn2PositiveIntAndPositiveLimit(t *testing.T) {
	int1 := 2
	int2 := 5
	limit := 10

	str1 := "clair-de-"
	str2 := "lune"

	goal := fmt.Sprintf("1,%s,3,4,%s,6,7,8,9,%s%s", str1, str2, str1, str2)

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
