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
	fmt.Printf("at position 3 we have: %d\n", ll.Get(3))
	ll.Set(2, 1000)
	ll.Print()
	newList := linkedlist.New()
	newList.InsertAtHead(2)
	newList.InsertAtTail(3)
	ll.Concatenate(*newList)
	ll.Print()
	shallow := linkedlist.New()
	shallow.ShallowCopy(*ll)
	shallow.Print()
	deep := linkedlist.New()
	deep.DeepCopy(*ll)
	deep.Print()

	fmt.Println("For list ll")
	ll.PrintAddresses()
	fmt.Println("For list newList")
	newList.PrintAddresses()
	fmt.Println("For list shallow")
	shallow.PrintAddresses()
	fmt.Println("For list deep")
	deep.PrintAddresses()
}
