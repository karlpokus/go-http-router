package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func TestRouter(t *testing.T) {
	r := New()
	r.Handler("/", http.HandlerFunc(foo))

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := w.Result()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		t.Errorf("statuscode want %d, got %d", 200, res.StatusCode)
	}
	if string(body) != "hello" {
		t.Errorf("body want %s, got %s", "hello", string(body))
	}
}
