package main

import (
	"fmt"

	bs "github.com/abelsouzacosta/data-structures-with-go/src/algorithms/binary_search"
	linkedlist "github.com/abelsouzacosta/data-structures-with-go/src/data_structures/linked_list"
)

func main() {
	l := linkedlist.New()

	l.InsertAtHead(1)
	l.InsertAtHead(2)
	l.InsertAtHead(3)
	l.InsertAtHead(4)
	l.InsertAtHead(5)
	l.InsertAtHead(6)
	l.InsertAtTail(7)
	l.InsertAtTail(8)
	l.InsertAtTail(9)
	l.InsertAtHead(10)

	l.Print()

	numbers := []int{
		1, 3, 6, 10, 15, 21, 28, 36,
		45, 55, 66, 78, 91, 105, 120, 136,
		153, 171, 190, 210, 231, 253, 276, 300,
		325, 351, 378, 406, 435, 465, 496, 528,
	}

	fmt.Printf("Target %d was found at %d\n", 21, bs.BinarySearch(numbers, 21)) // 5
	fmt.Printf("Target %d was found at %d\n", 91, bs.BinarySearch(numbers, 91)) // 11
}
