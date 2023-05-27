package main

import "fmt"

type Stack struct {
	slice []int
}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() int {
	var ret int = s.Peek()
	s.slice = s.slice[:len(s.slice)-1]
	return ret
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func main() {
	var s *Stack = new(Stack)
	s.Push(3)
	s.Push(4)
	s.Push(8)
	s.Push(5)
	s.Push(9)
	s.Push(7)
	fmt.Println(s)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	s.Push(10)
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
}
