package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	first *Node
	last *Node
	length int
}

func (q *Queue) printQueue() {

	current := q.first

	for current != nil {

		fmt.Printf("%v ---> %v\n",current.data, current.next)
		current = current.next
	}
}

func (q *Queue) peek() {
	fmt.Println(q.first)
}

func (q *Queue) enqueue(value interface{}) {
	
	newNode := &Node{data:value}

	if q.length == 0 {
		q.first = newNode
		q.last = newNode
		q.length++
		return
	}

	q.last.next = newNode
	q.last = newNode
	q.length++
	
}

func (q *Queue) dequeue() *Node {

	if q.isEmpty() {
		fmt.Println("empty queue")
		return nil
	}

	if q.length == 1 {
		q.last = nil
	}

	
	firstQueue := q.first
	q.first = q.first.next
	q.length--
	
	return firstQueue
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

func main() {
	// Test out the queue here and
	testQueue := &Queue{}

	// Test input stored in a slice
	testInput := []string{"Google","Quora","Udemy"}

	// Enqueueing testInput elements in testQueue
	for _, v := range testInput {
		testQueue.enqueue(v)
	}

	// Printing queue from first to last
	testQueue.printQueue()
	
}