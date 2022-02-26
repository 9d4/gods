package maxheap

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

// Adds an element to the heap
func (h *MaxHeap) Insert(element int) {
	h.array = append(h.array, element)
	h.maxHeapifyUp(len(h.array) - 1)
}

// return largest element
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]

	length := len(h.array)

	// move last child to top
	h.array[0] = h.array[length-1]

	// shorten the array
	h.array = h.array[:length-2]

	// then do heapify down
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index) // swap them
		index = parent(index)        // update current index
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	last := len(h.array) - 1
	indexToCompare := 0
	l, r := left(index), right(index)

	// do loop if index has at least one child
	for l <= last {
		switch {
		// if child is the last (the only child)
		case l == last:
			indexToCompare = l

		// the left child is larger
		case h.array[l] > h.array[r]:
			indexToCompare = l

		// the right child is larger
		default:
			indexToCompare = r
		}

		if h.array[index] < h.array[indexToCompare] {
			h.swap(index, indexToCompare)
			index = indexToCompare
			l, r = left(index), right(index)
		}
	}

}

// swap value between index
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// get the parrent
func parent(i int) int {
	return (i - 1) / 2
}

// get left child
func left(i int) int {
	return i*2 + 1
}

// get right child
func right(i int) int {
	return i*2 + 2
}

// helper to make new MaxHeap
func NewMaxHeap(elements ...int) *MaxHeap {
	h := &MaxHeap{}

	for _, item := range elements {
		h.Insert(item)
	}

	return h
}

// execute from main
func HeapMax() {
	aHeap := NewMaxHeap(10, 20, 40, 30, 50, 5)
	fmt.Println("aHeap    :", aHeap)

	for i, item := range aHeap.array {
		fmt.Printf("index %d  : %d\n", i, item)
		fmt.Printf("parent %d : index %d\n", i, parent(i))

		fmt.Println()
	}

	fmt.Println("Extract  :", aHeap.Extract())
	fmt.Println("aHeap    :", aHeap)
	fmt.Println()

	for i, item := range aHeap.array {
		fmt.Printf("index %d  : %d\n", i, item)
		fmt.Printf("parent %d : index %d\n", i, parent(i))

		fmt.Println()
	}
}
