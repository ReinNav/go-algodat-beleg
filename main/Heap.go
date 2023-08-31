package main

type MinHeap struct {
	nodes []*Node
}

type MaxHeap struct {
	nodes []*Node
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MinHeap) Size() int {
	return len(h.nodes)
}

func (h *MaxHeap) Size() int {
	return len(h.nodes)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.nodes) == 0
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.nodes) == 0
}

func (h *MinHeap) Insert(node *Node) {
	h.nodes = append(h.nodes, node)
	h.bubbleUp(h.Size() - 1)
}

func (h *MaxHeap) Insert(node *Node) {
	h.nodes = append(h.nodes, node)
	h.bubbleUp(h.Size() - 1)
}

func (h *MinHeap) ExtractMin() *Node {
	if h.IsEmpty() {
		return nil
	}
	min := h.nodes[0]
	h.nodes[0] = h.nodes[h.Size()-1]
	h.nodes = h.nodes[:h.Size()-1]
	h.bubbleDown(0)
	return min
}

func (h *MinHeap) bubbleUp(i int) {
	parent := (i - 1) / 2
	if i > 0 && h.nodes[i].distanceFromStart < h.nodes[parent].distanceFromStart {
		h.nodes[i], h.nodes[parent] = h.nodes[parent], h.nodes[i]
		h.bubbleUp(parent)
	}
}

func (h *MinHeap) bubbleDown(i int) {
	// childern
	left := 2*i + 1
	right := 2*i + 2
	min := i

	// find the smallest of left, right, and i
	if left < h.Size() && h.nodes[left].distanceFromStart < h.nodes[i].distanceFromStart {
		min = left
	}
	if right < h.Size() && h.nodes[right].distanceFromStart < h.nodes[min].distanceFromStart {
		min = right
	}
	if min != i {
		h.nodes[i], h.nodes[min] = h.nodes[min], h.nodes[i]
		h.bubbleDown(min)
	}
}

func (h *MaxHeap) ExtractMax() *Node {
	if h.IsEmpty() {
		return nil
	}
	max := h.nodes[0]
	h.nodes[0] = h.nodes[h.Size()-1]
	h.nodes = h.nodes[:h.Size()-1]
	h.bubbleDown(0)
	return max
}

func (h *MaxHeap) bubbleUp(i int) {
	parent := (i - 1) / 2
	if i > 0 && h.nodes[i].distanceFromStart > h.nodes[parent].distanceFromStart {
		h.nodes[i], h.nodes[parent] = h.nodes[parent], h.nodes[i]
		h.bubbleUp(parent)
	}
}

func (h *MaxHeap) bubbleDown(i int) {
	// children
	left := 2*i + 1
	right := 2*i + 2
	max := i

	// find the largest of left, right, and i
	if left < h.Size() && h.nodes[left].distanceFromStart > h.nodes[i].distanceFromStart {
		max = left
	}
	if right < h.Size() && h.nodes[right].distanceFromStart > h.nodes[max].distanceFromStart {
		max = right
	}
	if max != i {
		h.nodes[i], h.nodes[max] = h.nodes[max], h.nodes[i]
		h.bubbleDown(max)
	}
}
