// Package linkedlist deals with singly linked lists
package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

func New() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

func (list *LinkedList) InsertAtHead(data int) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	oldHead := list.head
	node.next = oldHead
	list.head = node
}

func (list *LinkedList) InsertAtTail(data int) {
	node := &Node{data: data, next: nil}

	if !list.IsEmpty() {
		oldTail := list.tail
		oldTail.next = node
		list.tail = node
		return
	}

	list.tail = node
	list.head = node
}

func (list *LinkedList) Print() {
	if list.IsEmpty() {
		fmt.Print("empty list")
	}

	current := list.head
	for current != nil {
		if current.next == nil {
			fmt.Printf("%d\n", current.data)
		} else {
			fmt.Printf("%d -> ", current.data)
		}
		current = current.next
	}
}
