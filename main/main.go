package main

import (
	"LearnGo/dataStructures"
	"fmt"
)

func main() {
	s := dataStructures.CreateStack[int](2, true)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())

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
