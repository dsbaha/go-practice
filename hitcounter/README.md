# Description
This is just a golang version of a practice hitcounter.  This is a multi-threaded version.

# Testing
```console
$ go test
PASS
ok      github.com/dsbaha/go-practice/hitcounter        0.013s
```

# Benchmarking
```console
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/dsbaha/go-practice/hitcounter
cpu: Intel(R) Atom(TM) x5-Z8350  CPU @ 1.44GHz
BenchmarkWorkLoop-4      2909568               433.7 ns/op
PASS
ok      github.com/dsbaha/go-practice/hitcounter        1.706s
```
