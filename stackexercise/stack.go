package main

import (
	"fmt"
	"math/rand"
)

// Stack is ... description
type Stack struct {
	data  []int
	min   []*int // this will be a pointer to the int in data
	max   []*int // this will be a pointer to the int in data
	count int
	sum   int
}

// NewStack returns a pointer to  stack
func NewStack() *Stack {
	return &Stack{
		data:  []int{},
		min:   []*int{},
		max:   []*int{},
		count: 0,
		sum:   0,
	}
}

func (s *Stack) push(i int) {
	s.data = append(s.data, i)
	s.count++
	s.sum += i
	ref := &s.data[s.count-1]

	// if its the min
	if len((*s).min) == 0 || len((*s).max) == 0 {
		(*s).min = append((*s).min, ref)
		(*s).max = append((*s).max, ref)
	} else if i < *(*s).min[len((*s).min)-1] {
		(*s).min = append((*s).min, ref)
	} else if i > *(*s).max[len((*s).max)-1] {
		(*s).max = append((*s).max, ref)
	}
}

func (s *Stack) pop() int {
	// get last element
	last := (*s).data[s.count-1]
	if last == *s.min[len((*s).min)-1] {
		s.min = s.min[:len((*s).min)-1]
	} else if last == *s.max[len((*s).max)-1] {
		s.max = s.max[:len((*s).max)-1]
	}
	// remove the last element
	s.data = s.data[:s.count-1]
	s.count--
	s.sum -= last
	return last
}

func (s *Stack) avg() float64 {
	return float64(s.sum) / float64(s.count)
}

func (s *Stack) print() {
	fmt.Printf("min: %d, max: %d, sum: %d, count: %d, avg: %.2f\n", *s.min[len(s.min)-1], *s.max[len(s.max)-1],
		s.sum, s.count, s.avg())
}

func main() {
	s := NewStack()
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
