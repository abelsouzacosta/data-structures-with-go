// Package linkedlist deals with the implementation of the linked list data structure
package linkedlist

type Node struct {
	data int
	next *Node
}

// setData defines the data that the Node holds
func (node *Node) setData(data int) {
	node.data = data
}

