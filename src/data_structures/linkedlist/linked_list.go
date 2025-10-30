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

func (list *LinkedList) InsertAfterElement(data, reference int) {
	node := &Node{data: data}
	if list.head == nil {
		return
	}
	current := list.head
	for current != nil && current.data != reference {
		current = current.next // moves forward until reference
	}
	if current == nil {
		return // reference not found in the list
	}
	node.next = current.next
	current.next = node
	if current == list.tail {
		list.tail = node
	}
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

func (list *LinkedList) DeleteElement(reference int) {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		return
	}
	if list.head.data == reference {
		list.DeleteFromHead()
		return
	}
	if list.tail.data == reference {
		list.DeleteFromTail()
		return
	}
	current := list.head
	for current != nil && current.next != nil && current.next.data != reference {
		current = current.next
	}
	if current == nil || current.next == nil {
		return // reference not found in the list
	}
	// found node
	nodeToDelete := current.next
	current.next = nodeToDelete.next
}

func (list *LinkedList) ReverseList() {
	// the list is empty or unary
	if list.head == nil || list.head.next == nil {
		return
	}
	oldHead := list.head
	current := list.head
	var previous *Node = nil
	var next *Node = nil
	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}
	list.head = previous
	list.tail = oldHead
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
