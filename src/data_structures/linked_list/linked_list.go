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
	// node is created with the head as it next reference
	node := &Node{data: data, next: list.head}
	list.head = node
	// if the list is empty then also updates the tail
	if list.tail == nil {
		list.tail = node
	}
}

func (list *LinkedList) InsertAtTail(data int) {
	node := &Node{data: data}
	// if tail is nil, it means that the list is empty
	if list.tail == nil {
		list.InsertAtHead(data)
		return
	}
	list.tail.next = node
	list.tail = node
}

func (list *LinkedList) DeleteFromHead() {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		list.head = nil
		list.tail = nil
		return
	}
	list.head = list.head.next
}

func (list *LinkedList) DeleteFromTail() {
	// list is empty
	if list.tail == nil {
		return
	}
	// list is unary
	if list.head.next == nil {
		list.head = nil
		list.tail = nil
		return
	}
	current := list.head
	for current.next != list.tail {
		current = current.next // moves ahead in the list
	}
	// updates the next reference to a nil value
	current.next = nil
	list.tail = current
}

func (list *LinkedList) Print() {
	if list.head == nil {
		fmt.Printf("[]")
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
