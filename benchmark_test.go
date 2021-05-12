package prioqueue_test

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/fgrosse/prioqueue"
)

var randValues []float32

func init() {
	rng := rand.New(rand.NewSource(42))
	randValues = make([]float32, 10_000)
	for i := range randValues {
		randValues[i] = rng.Float32()
	}
}

// BenchmarkMaxHeap_Push1_Empty tests how fast a single push operation is if the
// queue is not preallocated.
func BenchmarkMaxHeap_Push1_Empty(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	var h prioqueue.MaxHeap
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		h.Push(i, values[i])
	}
}

// BenchmarkMaxHeap_Push1_Preallocate tests how fast a single push operation is
// if the queue is preallocated.
func BenchmarkMaxHeap_Push1_Preallocate(b *testing.B) {
	rand.Seed(42)
	values := make([]float32, b.N)
	for i := range values {
		values[i] = rand.Float32()
	}

	n := uint32(b.N)

	h := prioqueue.NewMaxHeap(len(values))
	b.ResetTimer()
	b.ReportAllocs()

	for i := uint32(0); i < n; i++ {
		h.Push(i, values[i])
	}
}

// BenchmarkMaxHeap_Push200_Empty tests how fast we can push 200 elements on the
// MaxHeap implementation if we did not preallocate the queue.
func BenchmarkMaxHeap_Push200_Empty(b *testing.B) {
	h := new(prioqueue.MaxHeap)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			h.Push(uint32(id), val)
		}
	}
}

// BenchmarkMaxHeap_Push200_Preallocate tests how fast we can push 200 elements
// on the MaxHeap implementation if we preallocate the queue.
func BenchmarkMaxHeap_Push200_Preallocate(b *testing.B) {
	h := prioqueue.NewMaxHeap(len(randValues))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			h.Push(uint32(id), val)
		}
	}
}

// BenchmarkMaxHeap_Pop tests how fast a single pop operation of the MaxHeap
// implementation is when operating on 10,000 random elements.
func BenchmarkMaxHeap_Pop(b *testing.B) {
	pq := prioqueue.NewMaxHeap(len(randValues))
	for i := 0; i < len(randValues); i++ {
		pq.Push(uint32(i), randValues[i])
	}

	b.ResetTimer()
	b.ReportAllocs()
	for pq.Len() > 0 {
		pq.Pop()
	}
}

// BenchmarkStdlib_Push1_Empty tests how fast a single push operation is if the
// queue is not preallocated.
func BenchmarkStdlib_Push1_Empty(b *testing.B) {
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

// BenchmarkStdlib_Push1_Preallocate tests how fast a single push operation is
// if the queue is preallocated.
func BenchmarkStdlib_Push1_Preallocate(b *testing.B) {
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

// BenchmarkStdlib_Push200_Empty tests how fast we can push 200 elements on the
// StdHeap implementation if we did not preallocate the queue.
func BenchmarkStdlib_Push200_Empty(b *testing.B) {
	h := new(StdHeap)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			heap.Push(h, &prioqueue.Item{
				ID:   uint32(id),
				Prio: val,
			})
		}
	}
}

// BenchmarkStdlib_Push200_Preallocate tests how fast we can push 200 elements
// on the StdHeap implementation if we preallocate the queue.
func BenchmarkStdlib_Push200_Preallocate(b *testing.B) {
	h := make(StdHeap, 0, len(randValues))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			heap.Push(&h, &prioqueue.Item{
				ID:   uint32(id),
				Prio: val,
			})
		}
	}
}

// BenchmarkStdlib_Push200_EmptyInit tests how fast we can push 200 elements on
// the StdHeap implementation if we did not preallocate the queue but we call
// Init only once.
func BenchmarkStdlib_Push200_EmptyInit(b *testing.B) {
	h := new(StdHeap)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			h.Push(&prioqueue.Item{
				ID:   uint32(id),
				Prio: val,
			})
		}
		heap.Init(h)
	}
}

// BenchmarkStdlib_Push200_PreallocateInit tests how fast we can push 200
// elements on the StdHeap implementation if we preallocate the queue and
// call Init only once.
func BenchmarkStdlib_Push200_PreallocateInit(b *testing.B) {
	h := make(StdHeap, 0, len(randValues))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for id, val := range randValues {
			h.Push(&prioqueue.Item{
				ID:   uint32(id),
				Prio: val,
			})
		}
		heap.Init(&h)
	}
}

// BenchmarkStdlibHeap_Pop tests how fast a single pop operation of the StdHeap
// implementation is when operating on 10,000 random elements.
func BenchmarkStdlibHeap_Pop(b *testing.B) {
	h := make(StdHeap, 0, len(randValues))
	for i := 0; i < len(randValues); i++ {
		item := &prioqueue.Item{ID: uint32(i), Prio: randValues[i]}
		heap.Push(&h, item)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for h.Len() > 0 {
		heap.Pop(&h)
	}
}
