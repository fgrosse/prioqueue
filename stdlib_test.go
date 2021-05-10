package prioqueue_test

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/fgrosse/prioqueue"
	"github.com/stretchr/testify/assert"
)

// StdHeap is the heap implementation using the container/heap package to
// provide a baseline for benchmarks
type StdHeap []*prioqueue.Item

func (h StdHeap) Len() int {
	return len(h)
}

func (h StdHeap) Less(i, j int) bool {
	return h[i].Prio < h[j].Prio
}

func (h StdHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StdHeap) Push(x interface{}) {
	*h = append(*h, x.(*prioqueue.Item))
}

func (h *StdHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestStdHeap(t *testing.T) {
	q := new(StdHeap)
	n := 10_000

	rng := rand.New(rand.NewSource(42))
	for i := 0; i < n; i++ {
		prio := rng.Float32()
		heap.Push(q, &prioqueue.Item{ID: uint32(i), Prio: prio})
	}

	assert.Equal(t, n, q.Len())

	var last float32
	for q.Len() > 0 {
		item := heap.Pop(q).(*prioqueue.Item)
		if last != 0 && item.Prio < last {
			t.Errorf("Incorrect order: last %.0f popped=%.0f", last, item.Prio)
		}
		last = item.Prio
	}
	assert.Equal(t, 0, q.Len())
}

func BenchmarkStdlib_PushEmpty(b *testing.B) {
	rng := rand.New(rand.NewSource(42))
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rng.Float32()
	}

	n := uint32(b.N)

	h := new(StdHeap)
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		item := &prioqueue.Item{ID: i, Prio: values[i]}
		heap.Push(h, item)
	}
}

func BenchmarkStdlib_PushPreallocate(b *testing.B) {
	rng := rand.New(rand.NewSource(42))
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rng.Float32()
	}

	n := uint32(b.N)

	h := make(StdHeap, 0, len(values))
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		item := &prioqueue.Item{ID: i, Prio: values[i]}
		heap.Push(&h, item)
	}
}

func BenchmarkStdlibHeap_Pop(b *testing.B) {
	rng := rand.New(rand.NewSource(42))
	data := make([]uint32, b.N)
	prio := make([]float32, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rng.Uint32()
		prio[i] = rng.Float32()
	}

	h := new(StdHeap)
	for i := 0; i < len(data); i++ {
		item := &prioqueue.Item{ID: data[i], Prio: prio[i]}
		heap.Push(h, item)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for h.Len() > 0 {
		heap.Pop(h)
	}
}
