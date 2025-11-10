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

func (list *LinkedList) ReverseFirst(positions uint8) {
	if positions == 0 || list.head == nil || list.head.next == nil {
		return
	}
	current := list.head
	blockTail := list.head
	var previous *Node = nil
	var next *Node = nil
	var accumulator uint8 = 0
	for current != nil && accumulator < positions {
		next = current.next
		current.next = previous
		previous = current
		current = next
		accumulator++
	}

	// Case of partial reversion of the list
	// is not necessary to update the list tail
	// then we set the next element of the tail
	// of the block reversed to be the next element in the
	// original remaining list
	if next != nil {
		blockTail.next = next
	} else {
		// full reversion of the list
		// we must update the tail of the list in order
		// to keep list's integrity
		list.tail = blockTail
		blockTail.next = nil
	}

	list.head = previous
}

func (list *LinkedList) Get(index int) int {
	if list.head == nil {
		return 0
	}
	if list.head.next == nil {
		return list.head.data
	}
	current := list.head
	counter := 0
	for current != nil {
		if counter == index {
			return current.data
		}
		current = current.next
		counter += 1
	}
	return 0 // not found in the list
}

func (list *LinkedList) Set(index, data int) {
	if list.head == nil {
		return
	}
	if list.head.next == nil {
		list.head.data = data
	}
	current := list.head
	counter := 0
	for current != nil {
		if counter == index {
			current.data = data
		}
		current = current.next
		counter += 1
	}
}

func (list *LinkedList) Concatenate(listToConcatenate LinkedList) {
	if list.head == nil {
		list.head = listToConcatenate.head
		list.tail = listToConcatenate.tail
		return
	}
	list.tail.next = listToConcatenate.head
	list.tail = listToConcatenate.tail
}

func (list *LinkedList) ShallowCopy(origin LinkedList) {
	if origin.head == nil {
		return
	}
	list.head = origin.head
	list.tail = origin.tail
}

func (list *LinkedList) DeepCopy(origin LinkedList) {
	if origin.head == nil {
		return
	}
	current := origin.head
	for current != nil {
		list.InsertAtTail(current.data)
		current = current.next
	}
}

func (list *LinkedList) RemoveDuplicates() {
	if list.head == nil || list.head.next == nil {
		return
	}
	// keep track of the previous node
	current := list.head
	for current != nil {
		runner := current
		// stops in the last node
		for runner.next != nil {
			if runner.next.data == current.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
		current = current.next
	}
}

func (list *LinkedList) PrintAddresses() {
	if list.head == nil {
		return
	}
	fmt.Println("=================")
	fmt.Println("Printing Indexes")
	fmt.Println("=================")
	current := list.head
	for current != nil {
		fmt.Printf("%p\n", current)
		current = current.next
	}
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
