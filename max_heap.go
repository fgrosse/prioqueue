package prioqueue

type MaxHeap struct {
	items []*Item
}

func NewMaxHeap(size int) *MaxHeap {
	return &MaxHeap{
		items: make([]*Item, 0, size),
	}
}

func (h *MaxHeap) Top() (uint32, float32) {
	i := h.TopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

func (h *MaxHeap) TopItem() *Item {
	if len(h.items) == 0 {
		return nil
	}
	return h.items[0]
}

func (h *MaxHeap) Len() int {
	return len(h.items)
}

func (h *MaxHeap) Reset() {
	h.items = h.items[0:0]
}

// Push the value item into the priority queue with provided priority.
func (h *MaxHeap) Push(id uint32, prio float32) {
	item := &Item{ID: id, Prio: prio}
	h.PushItem(item)
}

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

func (h *MaxHeap) Pop() (id uint32, priority float32) {
	i := h.PopItem()
	if i == nil {
		return 0, 0
	}

	return i.ID, i.Prio
}

func (h *MaxHeap) PopItem() *Item {
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
func (h *MaxHeap) sink() {
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
