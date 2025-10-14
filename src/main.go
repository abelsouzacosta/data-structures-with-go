package main

import (
	"fmt"

	"github.com/abelsouzacosta/data-structures-with-go/src/data_structures/stack"
)

func main() {
	s := stack.New()

	s.Push(2)
	s.Push(3)
	s.Push(15)
	s.Push(33)
	s.Push(30)

	top, errTop := s.Peek()
	max, errMax := s.Max()
	min, errMin := s.Min()

	if errTop == nil {
		fmt.Println("The top value is: ", top)
	} else {
		fmt.Println("error: ", errTop)
	}

	if errMax == nil {
		fmt.Println("The max value is: ", max)
	} else {
		fmt.Println("error", errMax)
	}

	if errMin == nil {
		fmt.Println("The min value is: ", min)
	} else {
		fmt.Println("error", errMin)
	}

	s.Print()
}
