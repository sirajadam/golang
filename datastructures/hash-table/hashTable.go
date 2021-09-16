package main

import "fmt"

const arrSize = 7

type HashTable struct {
	data [arrSize]*linkedList
}

type linkedList struct {
	head   *node
	length int
}

type node struct {
	data int
	next *node
}


func (h *HashTable) Insert(key string) {
	
}
func (h *HashTable) Search(key string) bool {
	return false
}

func (h *HashTable) Delete(key string) {
	
}

func hash(key string) int {

	h := 0

	for pos, char := range key {
		h += pos * int(char)
	}

	return h % arrSize
}


func Init() *HashTable {
	newHashTable := &HashTable{}

	for i := range newHashTable.data {
		newHashTable.data[i] = &linkedList{}
	}

	return newHashTable
}

func main() {


	

	fmt.Println(hash("RANDY"))

}