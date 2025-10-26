package main

import (
	doublylinkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/doubly_linked_list"
)

func main() {
	dl := doublylinkedlist.New()

	dl.InsertAtHead(3)
	dl.InsertAtHead(4)
	dl.InsertAtTail(1)
	dl.InsertAfterElement(2, 3)
	dl.Print()
}
