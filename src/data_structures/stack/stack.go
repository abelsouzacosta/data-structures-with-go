// Package stack Holds the implementation of a Stack Data Structure
package stack

type Stack struct {
	elements []int
}

func New() *Stack {
	return &Stack{elements: []int{}}
}

// IsEmpty Checks if the stack is empty
// not allowing to remove elements from a empty stack
func (stack *Stack) IsEmpty() bool {
	return len(stack.elements) == 0
}

