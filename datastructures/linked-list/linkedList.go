package main

import (
	"fmt"
	"strings"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	data int
	next *Node
}

// printLinkedList prints the linkedlist in 
func (l *LinkedList) printLinkedList() {
	arr := []string{}
	current := l.head

	for current != nil {

		arr = append(arr, fmt.Sprint(current.data))
		current = current.next
	}

	str := strings.Join(arr, " --> ")
	fmt.Println(str)
	fmt.Printf("head: %v \t tail: %v \t length: %v\n", l.head.data, l.tail.data, l.length)
}

// prepend adds new value at the head of the linkedlist
func (l *LinkedList) prepend(val int) {

	newNode := &Node{data:val}

	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		newNode.next = l.head
		l.head = newNode
		l.length++
	}
}

// append adds new value at the tail of the linkedlist - O(1)
func (l *LinkedList) append(val int) {

	newNode := &Node{data: val}

	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
	} else {
		l.tail.next = newNode
		l.tail = newNode
		l.length++
	}
}

//traverseToindex traverses to index index in the linkedlist - O(n)
func (l *LinkedList) traverseToindex(index int) *Node{
	prevNode := l.head
	pos := 0

	for pos < index {
		prevNode = prevNode.next
		pos++
	}

	return prevNode
}

// insert inserts value at index index in linked list at index - O(n)
func (l *LinkedList) insert(index, value int) {

	if index < 0 || index >= l.length {
		fmt.Println("Index out of bounds")
		return
	}

	if index == l.length - 1 {
		l.append(value)
		return
	}
	
	if index == 0 {
		l.prepend(value)
		return
	}

	newNode := &Node{data:value}
	prevNode := l.traverseToindex(index-1)

	nextNodeHolder := prevNode.next
	prevNode.next = newNode
	newNode.next = nextNodeHolder
	l.length++
}

// remove removes the index in the linkedlist - O(n)
func (l *LinkedList) remove(index int) {

	if index < 0 || index >= l.length {
		fmt.Println("Index out of bounds")
		return
	}


	if index == 0 {
		l.head = l.head.next
		l.length--
		return
	}

	if index == l.length - 1 {
		prevNode := l.traverseToindex(index - 1)
		prevNode.next = nil
		l.tail = prevNode
		l.length--
		return
	}

	
	prevNode := l.traverseToindex(index - 1)
	prevNode.next = prevNode.next.next

	l.length--
}

// reverse reverses the linked list 
func (l *LinkedList) reverse() {

	if l.head.next == nil {
		return
	}

	first := l.head
	second := first.next
		 
	for second != nil {
		temp := second.next 
		second.next = first
		first = second 
		second = temp
	}

	l.head.next = nil
	l.head = first

	
} 

func main() {}