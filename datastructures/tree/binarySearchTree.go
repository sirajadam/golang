package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) InOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	bst.InOrderTraversal(root.left)
	fmt.Printf("%v ", root.value)
	bst.InOrderTraversal(root.right)
}

func (bst *BinarySearchTree) insert(val int) {

	newNode := &Node{value: val}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	currentNode := bst.root

	for {
		if val < currentNode.value {

			if currentNode.left == nil {
				currentNode.left = newNode
				return
			}
			currentNode = currentNode.left
		} else {

			if currentNode.right == nil {
				currentNode.right = newNode
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (bst *BinarySearchTree) lookup(val int)  {

	if bst.root == nil {
		return
	}

	currentNode := bst.root
	for currentNode != nil {
		if currentNode.value > val {
			currentNode = currentNode.left
		}  
		if currentNode.value < val {
			currentNode = currentNode.right
		} 
		if currentNode.value == val {
		
			fmt.Println("Found your value:", currentNode.value)
			return
			
		}
	}
	fmt.Println(2, "does not exist")
}

func main() {
	tree := &BinarySearchTree{}

	textInput := []int{9, 4, 6, 20, 170, 15, 1}

	for _, v := range textInput {
		tree.insert(v)
	}

//			9
//		4 		20
//	  1	  6  15    170

	tree.InOrderTraversal(tree.root)
	fmt.Println()
	tree.lookup(1)
}
