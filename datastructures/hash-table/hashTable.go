package main

import "fmt"

const arrSize = 7

// HashTable

type HashTable struct {
	data [arrSize]*linkedList
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.data[index].insert(key) 
}

func (h *HashTable) Search(key string) bool {
	return h.data[hash(key)].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.data[index].delete(key)
}

// LinkedList 

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data string
	next *node
}

func (list *linkedList) insert(key string) {
	if !list.search(key) {

		newNode := &node{key, nil}
		
		newNode.next = list.head
		list.head = newNode
		list.length++
		return
	} 
	fmt.Printf("%v -> already exists\n", key)
}

func (list *linkedList) search(key string) bool {
	current := list.head

	for current != nil {
		if current.data == key {
			return true
		}
		current = current.next
	}
	return false
}

func (list *linkedList) delete(key string) {

	// do nothing if list is empty
	if list.length == 0 {
		fmt.Println("list is empty")
		return
	} 
	
	// if key is head point next to head
	if list.head.data == key {
		list.head = list.head.next
		list.length--
		return
	}
	
	prevNode := list.head

	// keep track of previous node and compare current node
	for prevNode != nil {
		if prevNode.next.data == key {
			prevNode.next = prevNode.next.next
			return
		}
		prevNode = prevNode.next
	}

	fmt.Println("Element was not found : ", key)

}

func (list *linkedList) printLinkedListData() {

	current := list.head

	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}



func main() {


	h1 := Init()

	inputVal := []string{
		"oranges", 
		"apple",
		"banana",
		"sung jin-woo",
		"conor",
		"khabib",
		"tourist", 
	}

	for _, v := range inputVal{
		h1.Insert(v)
	}

	h1.Delete("banana")


	


}

func Init() *HashTable {
	newHashTable := &HashTable{}

	for i := range newHashTable.data {
		newHashTable.data[i] = &linkedList{}
	}

	return newHashTable
}



func hash(key string) int {

	h := 0

	for pos, char := range key {
		h += pos * int(char)
	}

	return h % arrSize
}
