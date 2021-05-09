package prioqueue_test

import (
	"math/rand"
	"testing"

	"github.com/fgrosse/prioqueue"
)

func TestMinHeap(t *testing.T) {
	var pq prioqueue.MinHeap
	runTests(t, &pq, assertSmallestFirst)

	pq2 := prioqueue.NewMinHeap(10)
	runTests(t, pq2, assertSmallestFirst)
}

func TestMinHeap_Random(t *testing.T) {
	pq := prioqueue.NewMinHeap(10)
	runTestsN(t, pq, assertSmallestFirst, 10_000)
}

func BenchmarkMinHeap_PushEmpty(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	var pq prioqueue.MinHeap
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		pq.Push(i, values[i])
	}
}

func BenchmarkMinHeap_PushPreallocate(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	pq := prioqueue.NewMinHeap(b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		pq.Push(i, values[i])
	}
}

func BenchmarkMinHeap_Pop(b *testing.B) {
	rand.Seed(42)
	data := make([]uint32, b.N)
	prio := make([]float32, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Uint32()
		prio[i] = rand.Float32()
	}

	pq := prioqueue.NewMinHeap(b.N)
	for i := 0; i < len(data); i++ {
		pq.Push(data[i], prio[i])
	}

	b.ResetTimer()
	b.ReportAllocs()
	for pq.Len() > 0 {
		pq.Pop()
	}
}
