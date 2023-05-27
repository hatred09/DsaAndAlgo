package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

func (q *Queue) Dequeue() int {
	var ret int = q.slice[0]
	q.slice = q.slice[1:]
	return ret
}

func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(3)
	fmt.Println(q)
	q.Enqueue(2)
	fmt.Println(q)
	q.Enqueue(5)
	fmt.Println(q)
	q.Enqueue(7)
	fmt.Println(q)
	q.Enqueue(4)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	q.Enqueue(8)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
