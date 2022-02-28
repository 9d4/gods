package minheap

import "fmt"

type MinHeap struct {
	array []int
}

func (h *MinHeap) Insert(element int) {
	h.array = append(h.array, element)
	h.minHeapifyUp(len(h.array) - 1)
}

func (h *MinHeap) Extract() int {
	// extracted item
	extracted := h.array[0]

	// move last index in to first index
	h.array[0] = h.array[len(h.array)-1]

	// shorthen the length of array
	h.array = h.array[:len(h.array)-2]

	// heapify down
	h.minHeapifyDown(0)

	// return
	return extracted
}

func (h *MinHeap) minHeapifyUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapifyDown(index int) {
	last := len(h.array) - 1
	indexToCompare := 0
	l, r := left(index), right(index)

	for last >= l {
		switch {
		case last == l:
			indexToCompare = l
		case h.array[l] < h.array[r]:
			indexToCompare = l
		case h.array[l] > h.array[r]:
			indexToCompare = r
		}

		if h.array[index] > h.array[indexToCompare] {
			h.swap(index, indexToCompare)
			index = indexToCompare
			l, r = left(index), right(index)
		}
	}
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

func NewMinHeap(elements ...int) *MinHeap {
	heap := &MinHeap{}

	for _, item := range elements {
		heap.Insert(item)
	}

	return heap
}

// execute from main
func HeapMin() {
	myHeap := NewMinHeap(50, 40, 20, 10, 30, 5)

	printHeap := func() {
		for i, item := range myHeap.array {
			fmt.Printf("index %d  : %d\n", i, item)
			fmt.Printf("parent %d : index %d\n", i, parent(i))

			fmt.Println()
		}
	}

	fmt.Println("INITIAL VALUES")
	fmt.Println("myHeap    :", myHeap)
	printHeap()

	fmt.Println("INSERTING 0")	
	myHeap.Insert(0)
	fmt.Println("myHeap    :", myHeap)
	printHeap()

	fmt.Println("Extract  :", myHeap.Extract())
	fmt.Println("myHeap    :", myHeap)
	printHeap()
}
