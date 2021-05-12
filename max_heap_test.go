package prioqueue_test

import (
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
