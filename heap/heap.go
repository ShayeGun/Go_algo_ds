package heap

import (
	"fmt"
)

type Heap struct{
	arr []int
}

func (h *Heap) BuildHeapify(arr []int) *Heap {
	h.arr = arr
	// index --> first non-leaf element of the tree
	for index:= len(h.arr) / 2 - 1; index >= 0; index--{
		h.heapify(h.arr, index)
	}
	
	return h
}

func (h *Heap) heapify(arr []int, index int) {
	max := index
	leftChild, rightChild := index * 2 + 1, index * 2 + 2
	length := len(arr)

	if (leftChild < length  && arr[leftChild] > arr[max]) {
		max = leftChild
	}

	if (rightChild < length  && arr[rightChild] > arr[max]) {
		max = rightChild
	}

	if (index != max){
		arr[index], arr[max] = arr[max], arr[index]
		h.heapify(arr, max)
	}
}

func (h *Heap) insert(num int) {
	h.arr = append(h.arr, num)
	lastIndex := len(h.arr) - 1
	h.heapUpWard(lastIndex)
}

func (h *Heap) heapUpWard(index int){
	// handle index out of range [-1] error
	if (index == 0) {
		return
	}

	parent := index / 2 - 1
	if (h.arr[parent] < h.arr[index] ){
		h.arr[parent], h.arr[index] = h.arr[index], h.arr[parent]
		h.heapUpWard(parent)
	}
}

func (h *Heap) delete() int {
	lastIndex := len(h.arr) - 1
	firstIndex := 0
	pop := h.arr[firstIndex]

	// swap first and last elements
	h.arr[lastIndex], h.arr[firstIndex] = h.arr[firstIndex], h.arr[lastIndex]
	// pop last element
	h.arr = h.arr[:len(h.arr) - 1]

	// rebuilding heap <only from root element>
	h.heapify(h.arr, firstIndex)

	return pop
}

func (h *Heap) sort() []int{
	length := len(h.arr)
	sortedArr := make([]int, length)
	for i := length - 1; i >= 0; i--{
		pop := h.delete()
		sortedArr[i] = pop
	}

	return sortedArr
}

func Printy(){
	h := Heap{arr: []int{}}
	h.BuildHeapify([]int{2,1,6,5,3,4,7})
	h.insert(8)
	fmt.Println(h.arr)
	sort := h.sort()
	fmt.Println(sort)
	fmt.Println(h.arr)
}