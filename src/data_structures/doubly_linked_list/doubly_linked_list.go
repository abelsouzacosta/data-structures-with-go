package doublylinkedlist

import "fmt"

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil}
}

func (list *DoublyLinkedList) isEmpty() bool {
	return list.tail == nil
}

func (list *DoublyLinkedList) insertAtEmptyList(data int) {
	node := &Node{data: data}
	list.head = node
	list.tail = node
}

func (list *DoublyLinkedList) InsertAtHead(data int) {
	node := &Node{previous: nil, next: nil, data: data}
	if list.isEmpty() {
		list.insertAtEmptyList(data)
		return
	}
	oldHead := list.head
	node.next = oldHead
	oldHead.previous = node
	list.head = node
}

func (list *DoublyLinkedList) InsertAtTail(data int) {
	node := &Node{data: data}
	if list.isEmpty() {
		list.insertAtEmptyList(data)
		return
	}
	oldTail := list.tail
	oldTail.next = node
	node.previous = oldTail
	list.tail = node
}

func (list *DoublyLinkedList) InsertAfterElement(data, reference int) {
	node := &Node{data: data}
	if list.isEmpty() {
		list.insertAtEmptyList(data)
		return
	}
	current := list.head
	for current != nil && current.data != reference {
		current = current.next
	}
	if current == nil {
		return // reference not found in the list
	}
	node.next = current.next
	current.next = node
	node.previous = current
	if node.next != nil {
		node.next.previous = node
	} else {
		list.tail = node
	}
}

func (list *DoublyLinkedList) DeleteAtHead() {
	if list.isEmpty() {
		return // cannot remove from an empty list
	}
	// list has only one element
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
		return
	}
	oldHead := list.head
	list.head = oldHead.next
	list.head.previous = nil
}

