package prioqueue_test

import (
	"math/rand"
	"testing"

	"github.com/fgrosse/prioqueue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type PriorityQueue interface {
	Push(id uint32, priority float32)
	Len() int
	Top() (id uint32, priority float32)
	Pop() (id uint32, priority float32)
	PopAndPush(*prioqueue.Item)
	Reset()
	Items() []*prioqueue.Item
}

type orderFunc func(current, last float32) bool

func assertSmallestFirst(current, last float32) bool {
	return last <= current
}

func assertBiggestFirst(current, last float32) bool {
	return last >= current
}

func runTests(t *testing.T, pq PriorityQueue, checkOrder orderFunc) {
	t.Helper()

	items := []prioqueue.Item{
		{ID: 1, Prio: 10},
		{ID: 2, Prio: 20},
		{ID: 3, Prio: 30},
		{ID: 4, Prio: 40},
		{ID: 5, Prio: 50},
		{ID: 6, Prio: 60},
		{ID: 7, Prio: 70},
		{ID: 8, Prio: 80},
		{ID: 9, Prio: 90},
		{ID: 10, Prio: 100},
	}

	rng := rand.New(rand.NewSource(42))
	for _, i := range rng.Perm(len(items)) {
		e := items[i]
		t.Logf("Adding item %+v", e)
		pq.Push(e.ID, e.Prio)
	}

	require.Equal(t, 10, pq.Len())

	pq.PopAndPush(&prioqueue.Item{ID: 11, Prio: 55})
	require.Equal(t, 10, pq.Len())

	t.Log("Item in array:")
	for _, item := range pq.Items() {
		t.Logf(" - id: %2.d prio: %3.0f", item.ID, item.Prio)
	}

	var last float32
	for pq.Len() > 0 {
		topID, topPrio := pq.Top()
		poppedID, poppedPrio := pq.Pop()
		t.Logf("Popped item %d: %.0f", poppedID, poppedPrio)
		assert.Equal(t, topID, poppedID)
		assert.Equal(t, topPrio, poppedPrio)
		if last != 0 && !checkOrder(poppedPrio, last) {
			t.Errorf("Incorrect order: last %.0f popped=%.0f", last, poppedPrio)
		}
		last = poppedPrio
	}

	pq.Reset()
	assert.Equal(t, 0, pq.Len())
}

func runTestsN(t *testing.T, pq PriorityQueue, checkOrder orderFunc, n int) {
	// Sanity checks on PriorityQueue to see it does not panic if it is empty.
	topID, topPrio := pq.Top()
	assert.EqualValues(t, 0, topID)
	assert.EqualValues(t, 0, topPrio)

	id, prio := pq.Pop()
	assert.EqualValues(t, 0, id)
	assert.EqualValues(t, 0, prio)

	rng := rand.New(rand.NewSource(42))
	for i := 0; i < n; i++ {
		prio := rng.Float32()
		pq.Push(uint32(i), prio)
	}

	assert.Equal(t, n, pq.Len())

	var last float32
	for pq.Len() > 0 {
		topID, topPrio := pq.Top()
		poppedID, poppedPrio := pq.Pop()
		assert.Equal(t, topID, poppedID)
		assert.Equal(t, topPrio, poppedPrio)
		if last != 0 && !checkOrder(poppedPrio, last) {
			t.Errorf("Incorrect order: last %.0f popped=%.0f", last, poppedPrio)
		}
		last = poppedPrio
	}
	assert.Equal(t, 0, pq.Len())
}
