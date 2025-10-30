package main

import linkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/linked_list"

func main() {
	ll := linkedlist.New()
	ll.InsertAtHead(3)
	ll.InsertAtHead(4)
	ll.InsertAtTail(5)
	ll.InsertAtTail(6)
	ll.InsertAfterElement(10, 5)
	ll.Print()
	ll.DeleteElement(3)
	ll.Print()
}
