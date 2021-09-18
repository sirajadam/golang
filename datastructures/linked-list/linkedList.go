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

func (l *LinkedList) traverseToindex(index int) *Node{
	prevNode := l.head
	pos := 0

	for pos < index {
		prevNode = prevNode.next
		pos++
	}

	return prevNode
}

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

func main() {

	myLinkedList := LinkedList{}


	// myLinkedList.prepend(5)
	// myLinkedList.prepend(4)
	// myLinkedList.prepend(3)
	// myLinkedList.prepend(2)
	// myLinkedList.prepend(1)
	myLinkedList.append(4)
	myLinkedList.append(5)
	myLinkedList.append(6)
	myLinkedList.prepend(2)
	myLinkedList.prepend(1)
	myLinkedList.prepend(0)
	myLinkedList.insert(3,3)
	myLinkedList.printLinkedList()
	myLinkedList.remove(3)
	fmt.Println("Removed index 3")
	myLinkedList.printLinkedList()
	myLinkedList.reverse()
	myLinkedList.printLinkedList()

}