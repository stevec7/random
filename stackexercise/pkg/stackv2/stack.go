package stackv2

import (
	"errors"
	"fmt"
)

// ErrStackEmpty is a custom error if the stack is empty
var ErrStackEmpty = errors.New("stack empty")

// MinMax min and max need to be computed in constant time, so
//	we need to keep a list of pointers to the previous val
type MinMax struct {
	node *Node
	next *Node
}

type Stack struct {
	top    *Node
	length int
	min    *MinMax
	max    *MinMax
	sum    int
}

func NewStack() *Stack {
	return &Stack{
		length: 0,
	}
}

type Node struct {
	next *Node
	data int
}

func (n *Node) Print() {
	fmt.Println(n.data)
}

func NewNode(d int) *Node {
	return &Node{
		data: d,
	}
}

func (s *Stack) Avg() float64 {
	return float64(s.sum) / float64(s.length)
}

func (s *Stack) Empty() bool {
	if s.Size() < 1 {
		return true
	}
	return false
}

func (s *Stack) Max() (*Node, error) {
	if s.max == nil {
		return &Node{}, ErrStackEmpty
	}
	return s.max.node, nil
}

func (s *Stack) Min() (*Node, error) {
	if s.min == nil {
		return &Node{}, ErrStackEmpty
	}
	return s.min.node, nil
}

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Push(n *Node) {
	n.next = s.top
	s.top = n
	s.sum += n.data
	s.length++

	// figure out min/max
	if s.min == nil {
		s.min = &MinMax{
			node: n,
			next: nil,
		}
	} else if n.data <= s.min.node.data {
		s.min = &MinMax{
			node: n,
			next: s.min.node,
		}
	}
	if s.max == nil {
		s.max = &MinMax{
			node: n,
			next: nil,
		}
	} else {
		if n.data >= s.max.node.data {
			s.max = &MinMax{
				node: n,
				next: s.max.node,
			}
		}
	}
}

func (s *Stack) Pop() (*Node, error) {
	if s.Size() < 1 {
		return &Node{}, ErrStackEmpty
	}
	n := s.top
	s.top = n.next
	s.sum = s.sum - n.data
	s.length--

	// now figure out if we need to remove from min/max
	if s.length > 1 {
		if n.data <= s.min.node.data {
			s.min = &MinMax{
				node: s.min.next,
				next: s.min.next.next,
			}
		} else if n.data >= s.max.node.data {
			s.max = &MinMax{
				node: s.max.next,
				next: s.max.next.next,
			}
		}
	} else if s.length == 1 {
		s.min = &MinMax{
			node: s.top,
			next: nil,
		}
		s.max = &MinMax{
			node: s.top,
			next: nil,
		}

	}
	return n, nil
}

func (s *Stack) Print() {
	if s.Empty() {
		fmt.Println("empty")
		return
	}
	n := s.top
	n.Print()
	newnode := n.next
	for i := 0; i < s.Size()-1; i++ {
		newnode.Print()
		newnode = newnode.next
	}
}

func (s *Stack) Sum() int {
	return s.sum
}
