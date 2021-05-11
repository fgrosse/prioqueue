# Go Priority Queue

Package `prioqueue` implements an efficient min and max priority queue using a
binary heap encoded in a slice.   

## Example usage

[embedmd]:# (example_test.go)
```go
package prioqueue_test

import (
	"fmt"
	"math/rand"

	"github.com/fgrosse/prioqueue"
)

func Example() {
	// n is the amount of elements we are about to push. If you don't know this
	// amount in advance then you can set n to 0 or a negative number. In this
	// case the slice uses by the queue will start with the default slice
	// capacity of Go.
	n := 10

	// We use a random number generator in this example to generate priority values.
	rng := rand.New(rand.NewSource(42))

	q := prioqueue.NewMaxHeap(n)
	for i := 0; i < n; i++ {
		// Every element we push and pop from the queue must have a unique identifier.
		// It is the callers responsibility to ensure this uniqueness.
		id := uint32(i)
		prio := rng.Float32()
		q.Push(id, prio)
	}

	// The queue will always return the highest priority element first.
	for q.Len() > 0 {
		id, prio := q.Pop()
		fmt.Printf("%.2f (id %d)\n", prio, id)
	}

	// Output:
	// 0.81 (id 6)
	// 0.65 (id 9)
	// 0.60 (id 2)
	// 0.38 (id 7)
	// 0.38 (id 5)
	// 0.38 (id 8)
	// 0.37 (id 0)
	// 0.21 (id 3)
	// 0.07 (id 1)
	// 0.04 (id 4)
}
```

## TODO: describe how it works

## TODO: explain why this is faster than the stdlib

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
