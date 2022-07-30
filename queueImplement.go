package main

import "fmt"

type Object interface {
	Enqueue()
	Top() int
	Dequeue() int
}
type Queue struct {
	queue []int
}

func (q *Queue) Enqueue(a int) {

	q.queue = append(q.queue, a)
}
func (q *Queue) Top() int {
	return q.queue[0]
}

func (q *Queue) Dequeue() int {
	temp := q.queue[0]
	q.queue = q.queue[1:]
	return temp
}

func main() {
	c := new(Queue)
	c.Enqueue(5)
	c.Enqueue(9)
	c.Enqueue(19)
	fmt.Println("Top:", c.Top())
	fmt.Println("Dequeue:", c.Dequeue())
	fmt.Println("Top:", c.Top())
	fmt.Println("Dequeue:", c.Dequeue())

}
