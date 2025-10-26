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

