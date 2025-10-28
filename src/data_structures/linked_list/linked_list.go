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

	if list.tail == nil {
		list.head = node
		list.tail = node
		return
	}

	oldTail := list.tail
	oldTail.next = node
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
	if list.tail == nil {
		return
	}

	if list.head.next == nil {
		list.head = nil
		list.tail = nil
		return
	}

	current := list.head
	for current.next != list.tail {
		current = current.next
	}
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
