package main

import (
	"fmt"

	"github.com/abelsouzacosta/data-structures-with-go/src/data_structures/stack"
)

func main() {
	fmt.Println("Hello World")
	s := stack.New()

	s.Push(2)
	s.Push(3)
	top, err := s.Peek()

	if err == nil {
		fmt.Println("The top value is: ", top)
	} else {
		fmt.Println("error: ", err)
	}

	s.Print()
}
