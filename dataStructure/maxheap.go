package dataStructure

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

//insert adds an element to the heap
func (maxHeap *MaxHeap) Insert(key int) {
	maxHeap.array = append(maxHeap.array, key)
	lastIndex := len(maxHeap.array) - 1
	maxHeap.maxHeapifyUpward(lastIndex)
}

func (maxHeap *MaxHeap) Extract() int {
	if len(maxHeap.array) <= 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	extractValue := maxHeap.array[0]

	lastIndex := len(maxHeap.array) - 1
	maxHeap.array[0] = maxHeap.array[lastIndex]
	maxHeap.array = maxHeap.array[:lastIndex]
	maxHeap.maxHeapifyDownward(0)

	return extractValue
}

//maxHeapifyUpward will heapify from bottom to top
func (maxHeap *MaxHeap) maxHeapifyUpward(index int) {
	for maxHeap.array[getParentIndex(index)] < maxHeap.array[index] {
		maxHeap.swapKeys(index, getParentIndex(index))
		//set index of item to parent since it swapped
		index = getParentIndex(index)
	}
}

func (maxHeap *MaxHeap) maxHeapifyDownward(index int) {
	//get the largerChild
	//compare to the current index
	//swap it  until it has no children
	lastIndex := len(maxHeap.array) - 1
	leftIndex, rightIndex := getLeftIndex(index), getRightIndex(index)
	childToCompare := 0
	for leftIndex <= lastIndex {
		childToCompare = maxHeap.getChildToCompare(leftIndex, rightIndex, lastIndex)
		if maxHeap.array[index] < maxHeap.array[childToCompare] {
			maxHeap.swapKeys(index, childToCompare)
			//update variables in the loop
			index = childToCompare
			leftIndex, rightIndex = getLeftIndex(index), getRightIndex(index)
		} else {
			return
		}
	}
}

func (maxHeap *MaxHeap) getChildToCompare(leftIndex, rightIndex, lastIndex int) int {
	// when left child is the only child i.e last index
	//when left child is larger
	if leftIndex == lastIndex || maxHeap.array[leftIndex] > maxHeap.array[rightIndex] {
		return leftIndex
	} else {
		return rightIndex
	}
}
func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftIndex(index int) int {
	return (index * 2) + 1
}

func getRightIndex(index int) int {
	return (index * 2) + 2
}

func (maxHeap *MaxHeap) swapKeys(firstIndex, secondIndex int) {
	maxHeap.array[firstIndex], maxHeap.array[secondIndex] = maxHeap.array[secondIndex], maxHeap.array[firstIndex]
}
