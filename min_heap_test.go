package prioqueue_test

import (
	"testing"

	"github.com/fgrosse/prioqueue"
)

func TestMinHeap(t *testing.T) {
	var pq prioqueue.MinHeap
	runTests(t, &pq, assertSmallestFirst)
}

//
// func TestNewMinHeap(t *testing.T) {
// 	pq2 := prioqueue.NewMinHeap(10)
// 	runTests(t, pq2, assertSmallestFirst)
// }
//
// func TestMinHeap_Random(t *testing.T) {
// 	pq := prioqueue.NewMinHeap(10)
// 	runTestsN(t, pq, assertSmallestFirst, 10_000)
// }
