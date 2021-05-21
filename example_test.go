package prioqueue_test

import (
	"fmt"
	"math/rand"

	"github.com/fgrosse/prioqueue"
)

func Example() {
	// The queue will be backed by a slice and knowing its size ahead of time
	// avoids unnecessary allocations. If you don't know in advance how many
	// items you want to push then you can set n to 0 or a negative number
	// instead. In this case the slice used by the queue will start with the
	// default slice capacity of Go.
	n := 10

	// We use a random number generator in this example to generate values.
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
