package prioqueue

// MinHeap implements a binary heap which is efficiently encoded in a slice.
// Since the heap is maintained in form of a binary tree, it can be represented
// in the form on an array.
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
// The priority queue has the following properties:
//   - items with low distance are dequeued before elements with higher distance (closest first)
//   - the item at the root of the tree is the minimum among all items present
//     in the binary heap. The same property is recursively true for all nodes in binary tree.
//
// Time Complexity
//
//   Push and Pop take O(log N) and Top() happens in constant time.
type MinHeap struct {
	items []*Item
}

type Item struct {
	ID   uint32
	Prio float32
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		items: make([]*Item, 0, size),
	}
}

func (h *MinHeap) Top() (uint32, float32) {
	i := h.TopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

func (h *MinHeap) TopItem() *Item {
	if len(h.items) == 0 {
		return nil
	}
	return h.items[0]
}

func (h *MinHeap) Len() int {
	return len(h.items)
}

func (h *MinHeap) Reset() {
	h.items = h.items[0:0]
}

// Push the value item into the priority queue with provided priority.
func (h *MinHeap) Push(id uint32, d float32) {
	item := &Item{ID: id, Prio: d}
	h.PushItem(item)
}

func (h *MinHeap) PushItem(item *Item) {
	// Add new item to the end of the list and then let it bubble up the binary
	// tree until the heap property is restored.
	h.items = append(h.items, item)

	i := len(h.items) - 1 // start at the last element
	for i > 0 {
		parent := (i - 1) / 2
		if h.items[parent].Prio <= h.items[i].Prio {
			// heap property is now satisfied again
			return
		}

		h.items[i], h.items[parent] = h.items[parent], h.items[i]
		i = parent
	}
}

func (h *MinHeap) Pop() (id uint32, priority float32) {
	i := h.PopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

func (h *MinHeap) PopItem() *Item {
	if len(h.items) == 0 {
		return nil
	}

	root := h.items[0]

	// swap first and last element
	h.items[0], h.items[len(h.items)-1] = h.items[len(h.items)-1], h.items[0]

	// remove last element
	h.items = h.items[0 : len(h.items)-1]

	// restore heap property
	h.sink()

	return root
}

// sink restores the heap property by shifting down the root node in the binary
// tree until the heap property is satisfied.
func (h *MinHeap) sink() {
	maxIndex := len(h.items) - 1
	i := 0 // start at the root node
	for {
		j := 2*i + 1 // index of first child of i
		if j > maxIndex {
			break // item i has no children
		}

		if j < maxIndex && h.items[j].Prio > h.items[j+1].Prio {
			j++
		}

		if h.items[i].Prio <= h.items[j].Prio {
			// heap property is now satisfied again
			break
		}

		// swap parent and child
		h.items[i], h.items[j] = h.items[j], h.items[i]

		// continue at child node
		i = j
	}
}
