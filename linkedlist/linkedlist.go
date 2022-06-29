package linkedlist

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head   *node
	length int
}

func (l *linkedlist) Prepend(n *node) {
	prevHead := l.head
	l.head = n
	l.head.next = prevHead
	l.length ++;
}

func (l *linkedlist) DeleteByValue(val int) {
	test := l.head
	var prevToDel *node

	for i := l.length; i > 0; i-- {
		if test.data == val {
			if prevToDel != nil {
				prevToDel.next = test.next
			}

			// the test is in the first place
			if test.next != nil && prevToDel == nil {
				l.head = test.next
			}

			l.length--
		}

		prevToDel = test
		test = test.next
	}

	return
} 

func (l linkedlist) GetList() (list []int) {
	node := l.head
	for i := l.length; i > 0; i-- {
		list = append(list, node.data)
		node = node.next
	}

	return
}

func LinkedList() {
	ll := new(linkedlist)
	ll.Prepend(&node{data: 90})
	ll.Prepend(&node{data: 100})
	ll.Prepend(&node{data: 80})
	ll.Prepend(&node{data: 45})
	ll.Prepend(&node{data: 23})
	ll.Prepend(&node{data: 10})
	ll.Prepend(&node{data: 23})

	fmt.Println(ll.GetList())

	// test: delete single number
	ll.DeleteByValue(100)
	fmt.Print("Deleted 100 ")
	fmt.Println(ll.GetList())

	// test: how if there are same value
	ll.DeleteByValue(23)
	fmt.Print("Deleted 23 ")
	fmt.Println(ll.GetList())

	// test: how if there is nothing inside the list
	anotherLl := &linkedlist{}
	anotherLl.DeleteByValue(1000)
}
