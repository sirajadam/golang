package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Stack struct {
	top *Node
	bottom *Node
	length int
}

func (s *Stack) peek() *Node{
	return s.top
}

func (s *Stack) push(value interface{}) {

	newNode := &Node{value, nil}

	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
		s.length++
		return
	} 

	newNode.next = s.top
	s.top = newNode
	s.length++
}

func (s *Stack) printStack() {

	if s.isEmpty() {
		fmt.Println("stack is empty")
		return
	}

	current := s.top

	for current != nil {
		fmt.Printf("|----%v----|\n", current.data)
		current = current.next
	}

	fmt.Printf("top: %v \t bottom: %v \t length: %v\n", s.top.data, s.bottom.data, s.length)

}

func (s *Stack) pop() interface{} {

	if s.top == nil {
		return nil
	}

	if s.length == 1 {
		s.bottom = nil
	}

	first := s.top
	s.top = s.top.next
	s.length--

	return first.data
}

func (s *Stack) isEmpty() bool {
	return s.length == 0
}



func main() {

	// Test out the stack in main func
	testStack := &Stack{}

	// Test input stored in a slice of strings
	testInput := []string{"Google","Quora","Udemy", "Discord"}

	// Pushing testInput elements into Stack
	for _, v := range testInput {
		testStack.push(v)
	}

	// Printing queue from top to bottom
	testStack.printStack()
}