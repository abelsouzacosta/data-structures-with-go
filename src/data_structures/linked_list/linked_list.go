package linkedlist

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

// New constructor for the linked list
func New() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

// IsEmpty checks if the linked list has no elements
func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

/*
* InsertAtHead will insert a element in the head of the list
*
* Creates a node that will be initialized with the data provided
* and sets the next reference of the node the the current head
* since the list is created having nil as the default value
* for head if the list is empty then the first element inserted
* will have nil in the next reference
 */
func (list *LinkedList) InsertAtHead(data int) {
	node := &Node{data: data, next: list.head}
	list.head = node

	// if there's no tail, then the list is empty
	// and the first node is the last node
	if list.tail == nil {
		list.tail = node
	}
}

/*
* InsertAtTail will insert a element in the tail of the list
*
* Creates a node and checks if the list is not empty
* then if the list is not empty it captures the current tail
* into a variable called oldTail and then updates this old tail
* next reference with the new node created, then points the tail list
* to the new node
*
* If the list is empty then both tail and head should point to the
* new node
 */
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
