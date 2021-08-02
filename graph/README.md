# Description
This is just a quick program that I made to learn Graph Data Structures and use Depth First Search (DFS) searching the linked nodes.  Features are that this is fully thread-safe leveraging the sync package (Read-Write Mutex)

# Compile
```console
$ go build
```

# Tests
```console
$ go test
PASS
ok      github.com/dsbaha/go-practice/graph     0.007s
```
# Benchmark
```console
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dsbaha/go-practice/graph
cpu: Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
BenchmarkFullGraph-4       32574             40950 ns/op
PASS
ok      github.com/dsbaha/go-practice/graph     1.728s
```
