package LinkedList

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (list *LinkedList) Prepend(n *Node) {
	second := list.Head
	list.Head = n
	n.Next = second
	list.Length++
}

func (list *LinkedList) PrintLinkedListData() {
	current := list.Head

	for current != nil {
		fmt.Printf("%v, ", current.Data)
		current = current.Next
	}
	fmt.Println("")
}

func (list *LinkedList) DeleteIndex(index int) {
	if index >= list.Length {
		fmt.Printf("IndexOutofBounds: Length of array %v, Index was: %v\n", list.Length, index )
		return
	}

	if index == 0 {
		list.Head = list.Head.Next
		list.Length--
		return
	}

	prevDelete := list.Head

	for i := 0; i < index - 1; i++ {
		prevDelete = prevDelete.Next
	}
	prevDelete.Next = prevDelete.Next.Next
	list.Length--
}

// Deletes ONLY first element found with @value
func (list *LinkedList) DeleteVal(value int) {

	if list.Length == 0 {
		fmt.Println("Nothing to delete in Empty LinkedList")
		return
	}

	if list.Head.Data == value {
		list.Head = list.Head.Next
		list.Length--
		return
	}

	prevDelete := list.Head

	for prevDelete.Next.Data != value {
		if prevDelete.Next.Next == nil {
			fmt.Println("Could not find value", value)
			return
		}
		prevDelete = prevDelete.Next
	}
	prevDelete.Next = prevDelete.Next.Next
	list.Length--
}

// func main() {

// 	// init empty linkedlist
// 	listOne := LinkedList{}
// 	// create 5 nodes
// 	n5 := Node{5, nil}
// 	n4 := Node{4, nil}
// 	n3 := Node{3, nil}
// 	n2 := Node{2, nil}
// 	n1 := Node{1, nil}
// 	// prepend n1-n5
// 	listOne.prepend(&n5)
// 	listOne.prepend(&n4)
// 	listOne.prepend(&n3)
// 	listOne.prepend(&n2)
// 	listOne.prepend(&n1)

// 	listOne.printLinkedListData()

// 	// test out more methods

// }
