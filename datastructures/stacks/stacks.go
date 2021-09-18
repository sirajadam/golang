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

	myStack := Stack{}
	
	myStack.push("google")
	myStack.push("udemy")
	myStack.push("discord")

	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
	fmt.Println(myStack.pop())
	
	myStack.printStack()
	fmt.Println(myStack)

	fmt.Println("")
}