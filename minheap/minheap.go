package minheap

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(element int) {
}

func (h *MinHeap) MinHeapifyUp(index int) {

}

func (h *MinHeap) MinHeapifyDown(index int) {

}

func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return (index * 2) + 1
}

func right(index int) int {
	return (index * 2) + 2
}

// execute from main
func HeapMin() {

}
