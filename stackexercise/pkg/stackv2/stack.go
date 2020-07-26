package stackv2

import "fmt"

type Stack struct {
	top    *Node
	length int
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

func (s *Stack) Size() int {
	return s.length
}

func (s *Stack) Push(n *Node) {
	n.next = s.top
	s.top = n
	s.length++
}

func (s *Stack) Pop() (*Node, error) {
	if s.Size() < 1 {
		return &Node{}, fmt.Errorf("stack empty")
	}
	n := s.top
	s.top = n.next
	s.length--
	return n, nil
}

func (s *Stack) Print() {
	n := s.top
	n.Print()
	newnode := n.next
	for i := 0; i < s.Size()-1; i++ {
		newnode.Print()
		newnode = newnode.next
	}
}
