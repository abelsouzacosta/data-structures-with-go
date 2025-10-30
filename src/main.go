package main

import (
	"fmt"

	linkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/linkedlist"
)

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
	ll.ReverseList()
	ll.Print()
	fmt.Printf("at position 3 we have: %d", ll.Get(3))
}
