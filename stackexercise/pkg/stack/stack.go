package stack

import (
	"errors"
	"fmt"
)

// Stack is my terrible way to create a stack data structure
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

// Avg returns the average of the elements in the stack
func (s *Stack) Avg() float64 {
	return float64(s.sum) / float64(s.count)
}

// Length returns the length of the stack
func (s *Stack) Length() int {
	return s.count
}

// Max returns the minimum in the stack
func (s *Stack) Max() (int, error) {
	if len(s.max) < 1 {
		return 0, errors.New("stack is empty")
	}
	return *s.max[len(s.max)-1], nil
}

// Min returns the minimum in the stack
func (s *Stack) Min() (int, error) {
	if len(s.min) < 1 {
		return 0, errors.New("stack is empty")
	}
	return *s.min[len(s.min)-1], nil
}

// Sum returns the sum of all elements on the stack
func (s *Stack) Sum() int {
	return s.sum
}

// Push adds an element to the stack
func (s *Stack) Push(i int) {
	s.data = append(s.data, i)
	s.count++
	s.sum += i

	// this is a pointer to the location when we inserted,
	//	so that if we need to insert this into the min or max
	//	slices, we dont run into a situation where they could be huge
	ref := &s.data[s.count-1]

	// using len() on s.min/max is constant time, as slices in Go
	//	are actually structs with fields:
	//	https://golang.org/src/runtime/slice.go
	//
	// first case is there isnt a min/max
	if len((*s).min) == 0 || len((*s).max) == 0 {
		(*s).min = append((*s).min, ref)
		(*s).max = append((*s).max, ref)
	} else if i < *(*s).min[len((*s).min)-1] {
		(*s).min = append((*s).min, ref)
	} else if i > *(*s).max[len((*s).max)-1] {
		(*s).max = append((*s).max, ref)
	}
}

// Pop removes an element from the end of the stack slice and returns it
func (s *Stack) Pop() (int, error) {
	// get last element, or return an error if the stack is empty
	if s.count == 0 {
		return 0, errors.New("stack is empty")
	}
	last := (*s).data[s.count-1]

	// using len() on s.min/max is constant time, see more in Push fcuntion comment
	if last == *s.min[len((*s).min)-1] {
		s.min = s.min[:len((*s).min)-1]
	} else if last == *s.max[len((*s).max)-1] {
		s.max = s.max[:len((*s).max)-1]
	}
	// remove the last element
	s.data = s.data[:s.count-1]
	s.count--
	s.sum -= last
	return last, nil
}

// Print shows the important parts of the stack
func (s *Stack) Print() {
	fmt.Printf("min: %d, max: %d, sum: %d, count: %d, avg: %.2f\n", *s.min[len(s.min)-1], *s.max[len(s.max)-1],
		s.sum, s.count, s.Avg())
}
