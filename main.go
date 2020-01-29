package main

import "fmt"

type Node struct {
	value string
	next *Node
	prev *Node
}

type Queue struct {
	first *Node
	last *Node
	size int
}

func (queue *Queue) Enqueue(str string){
	queue.size++
	if queue.first == nil {
		node := Node{
			value: str,
			next:  nil,
			prev:  nil,
		}
		queue.first = &node
		queue.last = &node
		return
	}

	copyPtr := queue.first
	node := Node{
		value: str,
		next:  nil,
		prev:  nil,
	}
	queue.first = &node
	queue.first.next = copyPtr
	copyPtr.prev = queue.first
}

func (queue *Queue) Dequeue() (string, error) {
	
	if queue.first == nil {
		return "", fmt.Errorf("no elements in queue")
	}

	defer func() {queue.size--}()
	basket := queue.last.value
	if queue.size == 1 {
		queue.first = nil
		queue.last = nil
		return basket, nil
	}

	beforeLast := queue.last.prev
	beforeLast.next = nil
	queue.last = beforeLast
	return basket, nil
}

func (queue *Queue) First() (string, error) {
	if queue.first == nil {
		return "", fmt.Errorf("empty queue")
	}
	return queue.first.value, nil
}

func (queue *Queue) Last() (string, error) {
	if queue.last == nil {
		return "", fmt.Errorf("empty queue")
	}
	return queue.last.value, nil
}

func (queue *Queue) Len() int {
	return queue.size
}

func main() {

}
