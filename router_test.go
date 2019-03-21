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
	path   string
	status int
	body   string
}{
	{"/", 200, "foo"},
	{"/missing", 404, "Not Found\n"},
}

func TestRouter(t *testing.T) {
	var r = New()
	r.Handler("/", http.HandlerFunc(foo))

	for _, tt := range testTable {
		req := httptest.NewRequest("GET", tt.path, nil)
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
