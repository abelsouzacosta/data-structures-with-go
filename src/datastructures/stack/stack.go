// Package stack Holds the implementation of a Stack Data Structure
package stack

import (
	"errors"
	"fmt"
	"math"
)

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

// Push inserts an element at the top of the stack
func (stack *Stack) Push(element int) {
	stack.elements = append(stack.elements, element)
}

// Pop removes an element from the top of
// the stack
func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("cannot pop from empty stack")
	}

	top := stack.elements[len(stack.elements)-1]
	stack.elements = stack.elements[:len(stack.elements)-1]
	return top, nil
}

// Peek returns the element at the top of the stack
// without removing it
func (stack *Stack) Peek() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("cannot peek from empty stack")
	}
	top := stack.elements[len(stack.elements)-1]
	return top, nil
}

func (stack *Stack) Print() {
	if stack.IsEmpty() {
		fmt.Println("empty stack")
	}

	for _, element := range stack.elements {
		fmt.Printf("%d\n", element)
	}
}

// Max return the maximum value within the stack
func (stack *Stack) Max() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("cannot get max element from empty stack")
	}

	max := 0

	for _, element := range stack.elements {
		if element > max {
			max = element
		}
	}

	return max, nil
}

// Min return the minimum value within the stack
func (stack *Stack) Min() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("cannot get min from empty stack")
	}

	min := math.MaxInt

	for _, element := range stack.elements {
		if element < min {
			min = element
		}
	}

	return min, nil
}
