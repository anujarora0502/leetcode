package solutions

import "fmt"

type PriorityQueue struct {
	queue []int
}

func Initialize() *PriorityQueue {
	return &PriorityQueue{queue: make([]int, 0)}
}

func (pq *PriorityQueue) Add(val int) {
	pq.queue = append(pq.queue, val)

	currIndex := len(pq.queue) - 1
	parentIndex := (currIndex - 1) / 2

	for parentIndex >= 0 && pq.queue[parentIndex] > pq.queue[currIndex] {
		temp := pq.queue[parentIndex]
		pq.queue[parentIndex] = pq.queue[currIndex]
		pq.queue[currIndex] = temp

		currIndex = parentIndex
		parentIndex = (currIndex - 1) / 2
	}
}

func (pq *PriorityQueue) Remove() int {
	priorityElement := pq.queue[0]
	pq.queue[0] = pq.queue[len(pq.queue)-1]
	pq.queue = pq.queue[:len(pq.queue)-1]

	currIndex := 0
	leftChildIndex := 2*currIndex + 1
	rightChildIndex := 2*currIndex + 2

	for leftChildIndex < len(pq.queue) && rightChildIndex < len(pq.queue) {
		if pq.queue[leftChildIndex] < pq.queue[currIndex] && pq.queue[leftChildIndex] < pq.queue[rightChildIndex]{
			temp := pq.queue[leftChildIndex]
			pq.queue[leftChildIndex] = pq.queue[currIndex]
			pq.queue[currIndex] = temp

			currIndex = leftChildIndex
			leftChildIndex = 2*currIndex + 1
			rightChildIndex = 2*currIndex + 2
		} else if pq.queue[rightChildIndex] < pq.queue[currIndex] {
			temp := pq.queue[rightChildIndex]
			pq.queue[rightChildIndex] = pq.queue[currIndex]
			pq.queue[currIndex] = temp

			currIndex = rightChildIndex
			leftChildIndex = 2*currIndex + 1
			rightChildIndex = 2*currIndex + 2
		} else {
			break
		}
	}

	if leftChildIndex < len(pq.queue) && pq.queue[leftChildIndex] < pq.queue[currIndex]{
		temp := pq.queue[leftChildIndex]
		pq.queue[leftChildIndex] = pq.queue[currIndex]
		pq.queue[currIndex] = temp
	}

	return priorityElement
}

func (pq *PriorityQueue) Size() int {
	return len(pq.queue)
}

func (pq *PriorityQueue) Peek() int {
	return pq.queue[0]
}

func (pq *PriorityQueue) PrintQueue() {
	fmt.Println(pq.queue)
}
