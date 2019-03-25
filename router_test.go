package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karlpokus/go-http-router/testdata/static"
	"github.com/karlpokus/go-http-router/testdata/mock"
)

var testTable = []struct {
	method string
	path   string
	status int
	body   string
}{
	{"GET", "/", 200, "foo"},
	{"POST", "/", 405, "Method Not Allowed\n"},
	{"GET", "/missing", 404, "Not Found\n"},
}

func TestRouter(t *testing.T) {
	rtr := New()
	rtr.Handler("GET", "/", http.HandlerFunc(mock.Foo))

	for _, tt := range testTable {
		r := httptest.NewRequest(tt.method, tt.path, nil)
		w := httptest.NewRecorder()
		rtr.ServeHTTP(w, r)
		res := w.Result()
		body, _ := ioutil.ReadAll(res.Body)

		if res.StatusCode != tt.status {
			t.Errorf("statuscode err: want %d, got %d", tt.status, res.StatusCode)
		}
		if string(body) != tt.body {
			t.Errorf("body err: want %s, got %s", tt.body, string(body))
		}
	}
}

func BenchmarkRouter(b *testing.B) {
	rtr := New()
	for _, r := range static.Routes {
		rtr.Handler(r.Method, r.Path, http.HandlerFunc(mock.Noop))
	}

	w := new(mock.ResponseWriter)
	r := httptest.NewRequest("GET", "/", nil)
	u := r.URL
	rq := u.RawQuery

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, staticroute := range static.Routes {
			r.Method = staticroute.Method
			r.RequestURI = staticroute.Path
			u.Path = staticroute.Path
			u.RawQuery = rq

			rtr.ServeHTTP(w, r)
		}
	}
}
