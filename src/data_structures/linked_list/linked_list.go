type LinkedList struct {
	head *Node
	tail *Node
}

// New constructor for the linked list
func New() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}
