package main

import (
	"LearnGo/dataStructures"
	"fmt"
	"log"
)

func main() {
	s := dataStructures.CreateStack[int](2, true)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	r1, e1 := s.Pop()
	r2, e2 := s.Pop()
	r3, e3 := s.Pop()
	r4, e4 := s.Pop()
	if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
		log.Fatal("Fatal!")
	}
	fmt.Println(r1, r2, r3, r4)
}

func multipleResultTest(x, y int) (a, b, c, d int) {
	a = x + y
	b = x - y
	c = x / y
	d = x * y
	return
}

type Some struct {
	i int
}
