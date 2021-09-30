package router

import (
	"github.com/monkeydioude/lbc-test/pkg/response"

	"net/http"
)

// Router's struct is made to implement http.Handler's
// interface while providing a simple "routing" API
// using the flexibility of http.HandlerFunc type
type Router struct {
	handlers map[string]map[string]http.HandlerFunc
}

// New returns a instance of a Router struct.
// Even so this will most likely be stored on the heap because of the map,
// considering the fact that this struct is going to be used almost always
// in main func, there is no reason to return a pointer.
func New() Router {
	return Router{
		handlers: make(map[string]map[string]http.HandlerFunc),
	}
}

// AddRoute binds a http.HandlerFunc to a HTTP method and a path.
func (r Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	if _, ok := r.handlers[method]; !ok {
		r.handlers[method] = make(map[string]http.HandlerFunc)
	}
	r.handlers[method][path] = handler
}

// Get is a wrapper around AddRoute forcing a GET HTTP method onto a path
// and a http.HandlerFunc
func (r Router) Get(path string, handler http.HandlerFunc) {
	r.AddRoute("GET", path, handler)
}

// Post is a wrapper around AddRoute forcing a POST HTTP method onto a path
// and a http.HandlerFunc
func (r Router) Post(path string, handler http.HandlerFunc) {
	r.AddRoute("POST", path, handler)
}

// ServeHTTP implements http.Handler interface while being the core
// this "routing" API
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if _, ok := r.handlers[req.Method]; !ok {
		response.NotFound(w)
		return
	}

	if _, ok := r.handlers[req.Method][req.URL.Path]; !ok {
		response.NotFound(w)
		return
	}

	r.handlers[req.Method][req.URL.Path](w, req)
}
