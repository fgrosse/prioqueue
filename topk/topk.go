// Package topk selects the items with the lowest or highest priority from a
// larger set of items using a bounded priority queue. This needs O(n log k)
// time and O(k) memory instead of sorting all n items.
package topk

import "github.com/fgrosse/prioqueue"

// Smallest returns the k items with the lowest priority, ordered from the
// lowest to the highest priority.
func Smallest(items []prioqueue.Item, k int) []prioqueue.Item {
	if k <= 0 {
		return nil
	}

	// Keep the k smallest items in a max heap so the largest of them can be
	// replaced whenever a smaller item is found.
	h := prioqueue.NewMaxHeap(k)
	for _, item := range items {
		switch {
		case h.Len() < k:
			h.Push(item.ID, item.Prio)
		case item.Prio < h.TopItem().Prio:
			h.PopAndPush(&prioqueue.Item{ID: item.ID, Prio: item.Prio})
		}
	}

	result := make([]prioqueue.Item, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		id, prio := h.Pop()
		result[i] = prioqueue.Item{ID: id, Prio: prio}
	}

	return result
}

// Largest returns the k items with the highest priority, ordered from the
// highest to the lowest priority.
func Largest(items []prioqueue.Item, k int) []prioqueue.Item {
	if k <= 0 {
		return nil
	}

	// Keep the k largest items in a min heap so the smallest of them can be
	// replaced whenever a larger item is found.
	h := prioqueue.NewMinHeap(k)
	for _, item := range items {
		switch {
		case h.Len() < k:
			h.Push(item.ID, item.Prio)
		case item.Prio > h.TopItem().Prio:
			h.PopAndPush(&prioqueue.Item{ID: item.ID, Prio: item.Prio})
		}
	}

	result := make([]prioqueue.Item, h.Len())
	for i := len(result) - 1; i >= 0; i-- {
		id, prio := h.Pop()
		result[i] = prioqueue.Item{ID: id, Prio: prio}
	}

	return result
}
