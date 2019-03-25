package router

import (
	"net/http"
)

type routes map[string][]route

type router struct {
	notFound         http.Handler
	methodNotAllowed http.Handler
	routes
}

type route struct {
	path    string
	handler http.Handler
}

// ServeHTTP calls the matched handler
func (rtr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.find(r.Method, r.URL.Path).ServeHTTP(w, r)
}

// find returns a handler matching the method and path
// returns a default error handler if no match
func (rtr *router) find(method, path string) http.Handler {
	methKey, ok := rtr.routes[method]
	if !ok {
		return rtr.methodNotAllowed
	}
	for _, r := range methKey {
		if r.path == path {
			return r.handler
		}
	}
	return rtr.notFound
}

// Handler adds a http.Handler to a http method and path
func (rtr *router) Handler(method, path string, handler http.Handler) {
	rtr.routes[method] = append(rtr.routes[method], route{path, handler})
}

// New returns a router
func New() *router {
	return &router{
		notFound:         http.HandlerFunc(notFound),
		methodNotAllowed: http.HandlerFunc(methodNotAllowed),
		routes:           make(routes),
	}
}

// notFound is the default 404 response
func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(404), 404)
}

// methodNotAllowed is the default 405 response
func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(405), 405)
}
