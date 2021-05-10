# Go Priority Queue

Package `prioqueue` implements an efficient min and max priority queue using a
binary heap encoded in a slice.   

# Example usage

TODO

# TODO: describe how it works

# TODO: explain why this is faster than the stdlib

## Benchmarks

This package provides a couple of benchmarks to understand the performance of
the implementation. In order to compare it with a baseline, there is also a
benchmark which uses the `container/heap` package from the standard library.

We can see that `github.com/fgrosse/prioqueue` has a slight edge over the
`container/heap` implementation when popping elements (`*Pop`), when pushing to
an empty queue (`*Empty`) or when pushing to a previously known amount of
elements onto the queue (`PushPreallocate`).

```shell
$ go test -run=None -bench 'Std|Max'
goos: linux
goarch: amd64
pkg: github.com/fgrosse/prioqueue
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkMaxHeap_PushEmpty-8         	12004537	        99.03 ns/op	      52 B/op	       1 allocs/op
BenchmarkMaxHeap_PushPreallocate-8   	33739531	        34.71 ns/op	       8 B/op	       1 allocs/op
BenchmarkMaxHeap_Pop-8               	 2418582	       804.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdlib_PushEmpty-8          	11076016	       119.4 ns/op	      55 B/op	       1 allocs/op
BenchmarkStdlib_PushPreallocate-8    	20723624	        54.81 ns/op	       8 B/op	       1 allocs/op
BenchmarkStdlibHeap_Pop-8            	 1547786	      1084 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/fgrosse/prioqueue	14.390s
```
