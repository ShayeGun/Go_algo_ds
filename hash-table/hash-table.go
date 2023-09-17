package hashTable

import (
	"fmt"
)

const ArraySize = 10

// make hash table struct
type HashTable struct {
	array [ArraySize]*node
}

// make linked list struct
type node struct {
	key string
	next *node
}

// initialize hash table
func (ht *HashTable) init () *HashTable{
	for i := range ht.array{
		ht.array[i] = &node{}
	}
	return ht
}

func (ht *HashTable) insertH(str string) *HashTable{
	index := hashing(str)
	newNode := node{str,ht.array[index]}
	ht.array[index] = &newNode
	return ht
}

func (ht *HashTable) searchH(str string) bool{
	index := hashing(str)
	node := ht.array[index]
	for node != nil{
		if node.key == str {
			return true
		}
		node = node.next
	}
	return false
}

func (ht *HashTable) deleteH(str string){
	index := hashing(str)
	preNode := ht.array[index]
	nextNode := preNode.next

	// if the first element must be deleted
	if preNode.key == str{
		ht.array[index] = nextNode
		preNode.next = nil
		return
	}
	
	// if other elements must be deleted
	for nextNode != nil{
		if nextNode.key == str {
			preNode.next = nextNode.next
			nextNode.next = nil
			return
		}
		preNode = nextNode
		nextNode = nextNode.next
	}
}

// hash function 
func hashing(str string) int{
	var sum int
	for _,v := range str{
		sum += int(v)
	}
	return sum % ArraySize
}

func Printy(){
	ht := HashTable{}
	ht.init()
	
	ht.insertH("ali")
	ht.insertH("reza")
	ht.insertH("al27")
	ht.deleteH("ali")
	fmt.Println(ht.array[0].next)

	fmt.Println(ht.searchH("ali"))
}