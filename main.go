package main

import (
	"fmt"

	"github.com/9d4/gods/array"
	"github.com/9d4/gods/maxheap"
	"github.com/9d4/gods/minheap"
)

func main() {
	fmt.Println("========================ARRAY========================")
	array.Array()
	fmt.Println("======================MAX HEAP=======================")
	maxheap.HeapMax()
	fmt.Println("======================MIN HEAP=======================")
	minheap.HeapMin()
}
