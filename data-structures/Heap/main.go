package main

import (
	"fmt"
)

type MinHeap struct {
	array []int
}

type MaxHeap struct {
	array []*int
}

func (min *MinHeap) insert(val int) {
	min.array = append(min.array, val)
	min.heapifyUpMin()
}

func (min *MinHeap) heapifyUpMin() {
	index := len(min.array) - 1
	for min.array[parentOf(index)] > min.array[index] {
		swap(min, index, parentOf(index))
		index = parentOf(index)
	}
}

func (min *MinHeap) minPop() int {
	popped := min.array[0]
	min.array[0] = min.array[len(min.array)-1]
	min.array = min.array[:len(min.array)-1]
	min.heapifyDownMin()
	return popped
}

func (min *MinHeap) heapifyDownMin() {
	index := 0
	smaller := leftChildOf(index)
	for leftChildOf(index) < len(min.array)-1 {
		if rightChildOf(index) < len(min.array) && min.array[rightChildOf(index)] < min.array[leftChildOf(index)] {
			smaller = rightChildOf(index)
		} else {
			smaller = leftChildOf(index)
		}
		if min.array[smaller] < min.array[index] {
			swap(min, smaller, index)
		}
		index = smaller
	}
}

func topKElements(nums []int, topElement int) []int {
	heap := MinHeap{}
	for i := 0; i < len(nums); i++ {
		heap.insert(nums[i])
		if i >= topElement {
			heap.minPop()

		}

	}
	return heap.array
}

// 1,2,3,4,5,6
func parentOf(val int) int {
	return (val - 1) / 2
}

func leftChildOf(val int) int {
	return (val * 2) + 1
}

func rightChildOf(val int) int {
	return (val * 2) + 2
}

func swap(heap *MinHeap, index1 int, index2 int) {
	heap.array[index1], heap.array[index2] = heap.array[index2], heap.array[index1]
}

func main() {
	heap := MinHeap{}

	heap.insert(5)
	heap.insert(4)
	heap.insert(3)
	heap.insert(2)
	heap.insert(1)
	fmt.Println(heap.array)
	heap.minPop()
	heap.minPop()
	heap.minPop()
	heap.minPop()
	fmt.Println(heap.array)
	heap.insert(4)
	heap.insert(30)
	heap.insert(25)
	heap.insert(11)
	heap.insert(41)
	heap.insert(320)
	heap.insert(255)
	heap.insert(111)
	fmt.Println(heap.array)
	heap.minPop()
	heap.minPop()
	heap.minPop()
	fmt.Println(heap.array)
	nums := []int{1, 2, 3, 4, 12, 34, 54, 123212, 34, 124, 311}
	slice := nums[:]
	fmt.Println(slice)
	fmt.Println(topKElements(slice, 3))
}
