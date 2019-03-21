package router

import (
	"net/http"
)

type router struct {
	routes []route
}

type route struct {
	Path    string
	Handler http.Handler
}

// ServeHTTP calls the matching handler for the url path
func (rtr *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rtr.find(r.URL.Path).ServeHTTP(w, r)
}

// find returns a handler matching the path
// if no match - returns a default 404 handler
func (rtr *router) find(path string) http.Handler {
	for _, r := range rtr.routes {
		if r.Path == path {
			return r.Handler
		}
	}
	return http.HandlerFunc(fourofour)
}

// Handler adds a http.Handler to a path
func (rtr *router) Handler(path string, handler http.Handler) {
	rtr.routes = append(rtr.routes, route{
		Path:    path,
		Handler: handler,
	})
}

// New returns a router
func New() *router {
	return &router{}
}

// fourofour is the default 404 response
func fourofour(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(404), 404)
}
