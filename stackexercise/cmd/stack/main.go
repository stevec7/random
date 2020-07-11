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
			s.push(-k)
		} else {
			s.push(k)
		}
	}
	s.print()
}
