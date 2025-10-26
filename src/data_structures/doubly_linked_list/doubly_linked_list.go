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

