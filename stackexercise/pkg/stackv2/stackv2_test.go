package stackv2

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	n1 := NewNode(5)
	n2 := NewNode(10)
	n3 := NewNode(3)
	s.Push(n1)
	s.Push(n2)
	s.Push(n3)
	s.Print()
	fmt.Printf("Size: %d\n", s.Size())
	fmt.Println("Popped")
	n, _ := s.Pop()
	fmt.Printf("Popped Node: \n")
	n.Print()
	fmt.Printf("NewStack: \n")
	s.Print()

}
