package doublylinkedlist

import "fmt"

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func New() *DoublyLinkedList {
	return &DoublyLinkedList{head: nil, tail: nil}
}

