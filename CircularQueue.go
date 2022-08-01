package main

import "fmt"

// declaring Queue with four fields a slice of queue, two pointers and one size of the queue
type Queue struct {
	queue []int
	front int
	rear  int
	size  int
}

// constructor for the queue where initialization of the field happens
// and return a queue
func NewCircularQueue() *Queue {
	return &Queue{
		queue: make([]int, 5),
		front: -1,
		rear:  -1,
		size:  5,
	}

}

// checks if the queue is full
func (q *Queue) IsFull() bool {
	// condition (circular) : if front is one step further than rear
	if q.front == q.rear+1 {
		return true
	}
	return false
}

// checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	// if q.front is at position -1 then its empty
	if q.front == -1 {
		return true
	}
	return false
}

// adding element to the queue
func (q *Queue) Enqueue(element int) {
	//first check if its full
	if q.IsFull() == true {
		fmt.Println("queue is full")
	} else if q.IsEmpty() == true {
		// if its empty(front = -1)move the front to 0
		q.front = 0
	}
	// else move the rear circularly and place the elements
	// in the rear postion
	q.rear = (q.rear + 1) % q.size
	q.queue[q.rear] = element
	fmt.Println("Inserted ", element)

}

func (q *Queue) Dequeue() int {
	var element int
	// check if empty
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
	} else if q.front == q.rear {
		//reset the queue
		q.front = -1
		q.rear = -1
	}
	// front value of the queue/ assign to the variable
	element = q.queue[q.front]
	// move the front circularly
	q.front = (q.front + 1) % q.size
	return element
}

func main() {
	obj := NewCircularQueue()
	obj.Enqueue(5)
	obj.Enqueue(9)
	obj.Enqueue(10)
	obj.Enqueue(11)
	obj.Enqueue(45)
	param_2 := obj.Dequeue()
	obj.Enqueue(9)
	param_3 := obj.IsFull()

	fmt.Println(param_2, param_3)
}
