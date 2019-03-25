package mock

import "net/http"

type ResponseWriter struct{}

func (rw *ResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (rw *ResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (rw *ResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (rw *ResponseWriter) WriteHeader(int) {}

func Foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}

func Noop(w http.ResponseWriter, r *http.Request) {}
