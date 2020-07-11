package main

import (
	"math/rand"

	"github.com/stevec7/random/stackexercise/pkg/stack"
)

func main() {
	s := stack.NewStack()
	for i := 0; i < 10; i++ {
		k := rand.Intn(16384)
		if i%7 == 0 {
			s.Push(-k)
		} else {
			s.Push(k)
		}
	}
	s.Print()
	for i := s
	_ = s.Pop()
	s.Print()
}
