package prioqueue_test

import (
	"math/rand"
	"testing"

	"github.com/fgrosse/prioqueue"
)

func TestMaxHeap(t *testing.T) {
	var pq prioqueue.MaxHeap
	runTests(t, &pq, assertBiggestFirst)

	pq2 := prioqueue.NewMaxHeap(10)
	runTests(t, pq2, assertBiggestFirst)
}

func TestMaxHeap_Random(t *testing.T) {
	pq := prioqueue.NewMaxHeap(10)
	runTestsN(t, pq, assertBiggestFirst, 10_000)
}

func BenchmarkMaxHeap_PushEmpty(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	var pq prioqueue.MaxHeap
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		pq.Push(i, values[i])
	}
}

func BenchmarkMaxHeap_PushPreallocate(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	pq := prioqueue.NewMaxHeap(len(values))
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		pq.Push(i, values[i])
	}
}

func BenchmarkMaxHeap_Pop(b *testing.B) {
	rand.Seed(42)
	data := make([]uint32, b.N)
	prio := make([]float32, b.N)
	for i := 0; i < len(data); i++ {
		data[i] = rand.Uint32()
		prio[i] = rand.Float32()
	}

	pq := prioqueue.NewMaxHeap(b.N)
	for i := 0; i < len(data); i++ {
		pq.Push(data[i], prio[i])
	}

	b.ResetTimer()
	b.ReportAllocs()
	for pq.Len() > 0 {
		pq.Pop()
	}
}
