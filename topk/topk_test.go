package topk

import (
	"testing"

	"github.com/fgrosse/prioqueue"
	"github.com/stretchr/testify/assert"
)

var items = []prioqueue.Item{
	{ID: 1, Prio: 5},
	{ID: 2, Prio: 1},
	{ID: 3, Prio: 8},
	{ID: 4, Prio: 3},
	{ID: 5, Prio: 9},
}

func TestSmallest(t *testing.T) {
	expected := []prioqueue.Item{{ID: 2, Prio: 1}, {ID: 4, Prio: 3}}
	assert.Equal(t, expected, Smallest(items, 2))
}

func TestLargest(t *testing.T) {
	expected := []prioqueue.Item{{ID: 5, Prio: 9}, {ID: 3, Prio: 8}, {ID: 1, Prio: 5}}
	assert.Equal(t, expected, Largest(items, 3))
}

func TestInvalidK(t *testing.T) {
	assert.Nil(t, Smallest(items, 0))
	assert.Nil(t, Largest(items, -1))
}
