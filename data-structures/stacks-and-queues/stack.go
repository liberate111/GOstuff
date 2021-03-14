package main

import "fmt"

//Stack represents a stack that hold a slice
type Stack struct {
	items []int
}

// Push will add a value at the end (top)
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end (top)
func (s *Stack) Pop() int {
	lastIdx := len(s.items) - 1
	toRemove := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
