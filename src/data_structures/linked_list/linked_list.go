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

