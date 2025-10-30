package main

import (
	"fmt"

	linkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/linkedlist"
)

func main() {
	ll := linkedlist.New()
	ll.InsertAtHead(10)
	ll.InsertAtHead(20)
	ll.InsertAtHead(30)
	ll.InsertAtHead(10)
	ll.InsertAtHead(30)
	ll.InsertAtHead(40)

	ll.Print()

	fmt.Println("Removing duplicates...")
	ll.RemoveDuplicates()
	ll.Print()
}
