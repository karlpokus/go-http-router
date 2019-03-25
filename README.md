# go-http-router
Benchmarking a custom http router in go. Let's see how fast I can make it compared to the [others](https://github.com/julienschmidt/go-http-routing-benchmark).

# usage
```bash
$ go test -bench=.

goos: darwin
goarch: amd64
pkg: github.com/karlpokus/go-http-router
BenchmarkRouter-4   	  100000	     15837 ns/op	       0 B/op	       0 allocs/op
```

I tried to reproduce the testing conditions in [the router benchmark repo](https://github.com/julienschmidt/go-http-routing-benchmark) but seeing as this beats the http.DefaultServeMux - on both speed and allocations - I think it's probably too good to be true. I'll investigate further later.

# todos
- [x] add method to route struct
- [x] benchmark static routes
- [ ] add url parameters

# license
MIT
