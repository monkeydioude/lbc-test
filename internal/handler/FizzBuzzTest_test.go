package handler

import (
	"net/http/httptest"
	"testing"
)

func TestICanReturnACorrectResponseOn2PositiveIntAndPositiveLimit(t *testing.T) {
	p := params{
		int1:  2,
		int2:  5,
		limit: 10,
		inter: 10,
		str1:  "clair-de-",
		str2:  "lune",
	}
	goal := "1,clair-de-,3,clair-de-,lune,clair-de-,7,clair-de-,9,clair-de-lune"
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestICanFailOnMissingParameters(t *testing.T) {
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?int1=1&str1=a&str2=b&limit=2", nil)
	_, err := buildParamsFromValues(reqtest.URL.Query())

	if err == nil {
		t.Fail()
	}
}

func TestICanFailOnUnwantedIntValues(t *testing.T) {
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?int1=1&int2=2&limit=a&str1=a&str2=b", nil)
	_, err := buildParamsFromValues(reqtest.URL.Query())

	if err == nil {
		t.Fail()
	}
}

func TestICanReturnNothingOnLimit0(t *testing.T) {
	p := params{
		int1:  2,
		int2:  5,
		limit: 0,
		inter: 10,
		str1:  "clair-de-",
		str2:  "vide",
	}
	if string(computeResponse(p)) != "" {
		t.Fail()
	}
}

func TestICanReturnNothingOnNegativeLimit(t *testing.T) {
	p := params{
		int1:  2,
		int2:  5,
		limit: -1,
		inter: 10,
		str1:  "clair-de-",
		str2:  "trou-noir",
	}
	if string(computeResponse(p)) != "" {
		t.Fail()
	}
}

func TestIDontDisplayInterIfInt2LargerThanLimit(t *testing.T) {
	goal := "1,int2?,3,int2?"

	p := params{
		int1:  2,
		int2:  5,
		limit: 4,
		inter: 10,
		str1:  "int2?",
		str2:  "*zZz*",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1EqualsInt2(t *testing.T) {
	goal := "1,woll-smoth,3,woll-smoth"

	p := params{
		int1:  2,
		int2:  2,
		limit: 4,
		inter: 2,
		str1:  "woll-",
		str2:  "smoth",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1IsNegative(t *testing.T) {
	goal := "1,:(,-why-so-negative-mate,:(,5,:(-why-so-negative-mate"

	p := params{
		int1:  -2,
		int2:  3,
		limit: 6,
		inter: -6,
		str1:  ":(",
		str2:  "-why-so-negative-mate",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestInt1AndInt2AreNegatives(t *testing.T) {
	goal := "1,together-we,-can-be-fine,together-we,5,together-we-can-be-fine"

	p := params{
		int1:  -2,
		int2:  -3,
		limit: 6,
		inter: 6,
		str1:  "together-we",
		str2:  "-can-be-fine",
	}
	if string(computeResponse(p)) != goal {
		t.Fail()
	}
}

func TestICanComputeInterIsEqualToInt1IfInt1EqualsInt2(t *testing.T) {
	// int1 and int2 must be same for this test
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/fizz-buzz/test?int1=2&int2=2&limit=2&str1=a&str2=b", nil)
	p, err := buildParamsFromValues(reqtest.URL.Query())

	if err != nil || p.limit != 2 {
		t.Fail()
	}
}
