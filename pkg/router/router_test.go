package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestICanMatchAPathAndAMethod(t *testing.T) {
	router := New()
	wtest := httptest.NewRecorder()
	reqtest := httptest.NewRequest("GET", "http://localhost:8004/salut", nil)
	router.Get("/salut", func(w http.ResponseWriter, req *http.Request) {})

	router.ServeHTTP(wtest, reqtest)
	if wtest.Result().StatusCode == 404 {
		t.Fail()
	}
}

func TestICanReturn404IfMethodDoesNotMatch(t *testing.T) {
	router := New()
	wtest := httptest.NewRecorder()
	// applying POST on request dummy
	reqtest := httptest.NewRequest("POST", "http://localhost:8004/salut", nil)
	// routing over GET
	router.Get("/salut", func(w http.ResponseWriter, req *http.Request) {
	})

	router.ServeHTTP(wtest, reqtest)
	if wtest.Result().StatusCode != 404 {
		t.Fail()
	}
}

func TestICanReturn404OIfPathDoesNotMatch(t *testing.T) {
	router := New()
	wtest := httptest.NewRecorder()
	// applying GET on request dummy but requesting /cpacool path
	reqtest := httptest.NewRequest("GET", "http://localhost:8004/cpacool", nil)
	// routing over GET but routing on /ccool
	router.Get("/ccool", func(w http.ResponseWriter, req *http.Request) {
	})

	router.ServeHTTP(wtest, reqtest)
	if wtest.Result().StatusCode != 404 {
		t.Fail()
	}
}
