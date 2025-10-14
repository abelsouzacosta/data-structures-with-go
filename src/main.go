package main

import linkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/linked_list"

func main() {
	l := linkedlist.New()

	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.InsertAtHead(4)
	l.InsertAtHead(5)
	l.InsertAtHead(6)

	l.Print()
}
