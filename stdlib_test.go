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

func (h *StdHeap) Len() int {
	return len(*h)
}

func (h *StdHeap) Less(i, j int) bool {
	return (*h)[i].Prio < (*h)[j].Prio
}

func (h *StdHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
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
