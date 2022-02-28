package main

import (
	"fmt"

	"github.com/traperwaze/gods/array"
	"github.com/traperwaze/gods/maxheap"
	"github.com/traperwaze/gods/minheap"
)

func main() {
	fmt.Println("========================ARRAY========================")
	array.Array()
	fmt.Println("======================MAX HEAP=======================")
	maxheap.HeapMax()
	fmt.Println("======================MIN HEAP=======================")
	minheap.HeapMin()
}
