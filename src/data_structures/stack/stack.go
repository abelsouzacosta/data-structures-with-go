// Package stack Holds the implementation of a Stack Data Structure
package stack

type Stack struct {
	elements []int
}

func New() *Stack {
	return &Stack{elements: []int{}}
}

