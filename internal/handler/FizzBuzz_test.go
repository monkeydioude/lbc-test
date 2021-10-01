package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/monkeydioude/lbc-test/internal/entity"
)

func TestICanReturnACorrectResponseOn2PositiveIntAndPositiveLimit(t *testing.T) {
	p := entity.Params{
		Int1:  2,
		Int2:  5,
		Limit: 10,
		Inter: 10,
		Str1:  "clair-de-",
		Str2:  "lune",
	}
	goal := "1,clair-de-,3,clair-de-,lune,clair-de-,7,clair-de-,9,clair-de-lune"
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestICanFailOnMissingParameters(t *testing.T) {
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?Int1=1&Str1=a&Str2=b&Limit=2", nil)
	_, err := buildParamsFromValues(reqtest.URL.Query())

	if err == nil {
		t.Fail()
	}
}

func TestICanFailOnUnwantedIntValues(t *testing.T) {
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?Int1=1&Int2=2&Limit=a&Str1=a&Str2=b", nil)
	_, err := buildParamsFromValues(reqtest.URL.Query())

	if err == nil {
		t.Fail()
	}
}

func TestICanReturnNothingOnLimit0(t *testing.T) {
	p := entity.Params{
		Int1:  2,
		Int2:  5,
		Limit: 0,
		Inter: 10,
		Str1:  "clair-de-",
		Str2:  "vide",
	}
	if string(computeResponse(p)) != "" {
		t.Fail()
	}
}

func TestICanReturnNothingOnNegativeLimit(t *testing.T) {
	p := entity.Params{
		Int1:  2,
		Int2:  5,
		Limit: -1,
		Inter: 10,
		Str1:  "clair-de-",
		Str2:  "trou-noir",
	}
	if string(computeResponse(p)) != "" {
		t.Fail()
	}
}

func TestIDontDisplayInterIfInt2LargerThanLimit(t *testing.T) {
	goal := "1,Int2?,3,Int2?"

	p := entity.Params{
		Int1:  2,
		Int2:  5,
		Limit: 4,
		Inter: 10,
		Str1:  "Int2?",
		Str2:  "*zZz*",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1EqualsInt2(t *testing.T) {
	goal := "1,woll-smoth,3,woll-smoth"

	p := entity.Params{
		Int1:  2,
		Int2:  2,
		Limit: 4,
		Inter: 2,
		Str1:  "woll-",
		Str2:  "smoth",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1IsNegative(t *testing.T) {
	goal := "1,:(,-why-so-negative-mate,:(,5,:(-why-so-negative-mate"

	p := entity.Params{
		Int1:  -2,
		Int2:  3,
		Limit: 6,
		Inter: -6,
		Str1:  ":(",
		Str2:  "-why-so-negative-mate",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1AndInt2AreNegatives(t *testing.T) {
	goal := "1,together-we,-can-be-fine,together-we,5,together-we-can-be-fine"

	p := entity.Params{
		Int1:  -2,
		Int2:  -3,
		Limit: 6,
		Inter: 6,
		Str1:  "together-we",
		Str2:  "-can-be-fine",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestICanComputeInterIsEqualToInt1IfInt1EqualsInt2(t *testing.T) {
	// Int1 and Int2 must be same for this test
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?Int1=2&Int2=2&Limit=2&Str1=a&Str2=b", nil)
	p, err := buildParamsFromValues(reqtest.URL.Query())

	if err != nil || p.Limit != 2 {
		t.Fail()
	}
}
