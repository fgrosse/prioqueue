// Package prioqueue implements an efficient min and max priority queue using a
// binary heap encoded in a slice.
package prioqueue

// MaxHeap implements a priority queue which allows to retrieve the highest
// priority element using a heap. Since the heap is maintained in form of a
// binary tree, it can efficiently be represented in the form of a list.
//
// The priority queue has the following properties:
//   - items with high priority are dequeued before elements with lower priority
//   - the item at the root of the tree is the maximum among all items present
//     in the binary heap. The same property is recursively true for all nodes
//     in the tree.
//
// Array representation
//
// The first element of the list is always the root node (R) of the binary tree.
// The two children of (R) are the next two elements in the list (A) & (B).
// (A) and (B) are immediately followed by the children of (A) and then the
// children of (B). This process continues for all nodes of the binary tree.
// Generally speaking, the parent of index i is at index (i-1)/2. The two
// children of index i are at (2*i)+1 and (2*i)+2.
//
// Time Complexity
//
//   Push and Pop take O(log n) and Top() happens in constant time.
type MaxHeap struct {
	items []*Item
}

// NewMaxHeap returns a new MaxHeap instance which contains a pre-allocated
// backing array for the stored items. Usage of this function or setting a
// correct size is optional. If more items are inserted into the queue than
// there is space in the backing array, it grows automatically.
func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		items: make([]*Item, 0, size),
	}
}

// Top returns the ID and priority of the item with the highest priority value
// in the queue without removing it.
func (h *MaxHeap) Top() (uint32, float32) {
	i := h.TopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

// TopItem returns the item with the highest priority value in the queue without
// removing it.
func (h *MaxHeap) TopItem() *Item {
	if len(h.items) == 0 {
		return nil
	}
	return h.items[0]
}

// Len returns the amount of elements in the queue.
func (h *MaxHeap) Len() int {
	return len(h.items)
}

// Reset is a fast way to empty the queue. Note that the underlying array will
// still be used by the heap which means that this function will not free up any
// memory. If you need to release memory, you have to create a new instance and
// let this one be taken care of by the garbage collection.
func (h *MaxHeap) Reset() {
	h.items = h.items[0:0]
}

// Push the value item into the priority queue with provided priority.
func (h *MaxHeap) Push(id uint32, prio float32) {
	item := &Item{ID: id, Prio: prio}
	h.PushItem(item)
}

// PushItem adds an Item to the queue.
func (h *MaxHeap) PushItem(item *Item) {
	// Add new item to the end of the list and then let it bubble up the binary
	// tree until the heap property is restored.
	h.items = append(h.items, item)

	i := len(h.items) - 1 // start at the last element
	for i > 0 {
		parent := (i - 1) / 2
		if h.items[parent].Prio >= h.items[i].Prio {
			// heap property is now satisfied again
			return
		}

		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

// Pop removes the item with the highest priority value from the queue and
// returns its ID and priority.
func (h *MaxHeap) Pop() (id uint32, priority float32) {
	i := h.PopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

// PopItem removes the item with the highest priority value from the queue.
func (h *MaxHeap) PopItem() *Item {
	if len(h.items) == 0 {
		return nil
	}

	root := h.items[0]
	maxIndex := len(h.items) - 1

	// swap first and last element and then remove the last from the list
	h.items[0], h.items[maxIndex] = h.items[maxIndex], h.items[0]
	h.items = h.items[0:maxIndex]

	// restore heap property
	h.shiftDown()

	return root
}

// shiftDown restores the heap property by shifting down the root node in the binary
// tree until the heap property is satisfied.
func (h *MaxHeap) shiftDown() {
	maxIndex := len(h.items) - 1
	i := 0 // start at the root node
	for {
		j := 2*i + 1 // index of first child of i
		if j > maxIndex {
			break // item i has no children
		}

		if j < maxIndex && h.items[j].Prio < h.items[j+1].Prio {
			j++
		}

		if h.items[i].Prio >= h.items[j].Prio {
			// heap property is now satisfied again
			break
		}

		// swap parent and child
		h.items[i], h.items[j] = h.items[j], h.items[i]

		// continue at child node
		i = j
	}
}
