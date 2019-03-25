package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

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
	r := New()
	r.Handler("GET", "/", http.HandlerFunc(foo))

	for _, tt := range testTable {
		req := httptest.NewRequest(tt.method, tt.path, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
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

var benchRes http.Handler // avoiding compiler optimisations

func BenchmarkRouter(b *testing.B) {
	var res http.Handler
	r := New()
	foo := http.HandlerFunc(foo)
	r.Handler("GET", "/", foo)
	for n := 0; n < b.N; n++ {
		//r.ServeHTTP(w, req)
		res = r.find("GET", "/missing")
	}
	benchRes = res
}
