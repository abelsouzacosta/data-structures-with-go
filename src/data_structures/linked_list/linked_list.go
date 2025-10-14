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
