package linkedlistexercise

import (
	"fmt"
	"strconv"
)

type Node struct {
	data int
	next *Node
}

func (n Node) Data() int {
	return n.data
}

func (n Node) String() string {
	return strconv.Itoa(n.data)
}

func NewNode(d int) *Node {
	return &Node{
		data: d,
		next: nil,
	}
}

type LinkedList struct {
	head *Node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList) Add(n *Node) {
	if l.head == nil {
		l.head = n
		n.next = nil
		l.size++
		return
	}

	// linked list isnt empty, add to the "tail"
	node := l.head
	for {
		if node.next != nil {
			node = node.next
		} else {
			node.next = n
			l.size++
			return
		}
	}
}

func (l *LinkedList) Empty() bool {
	return l.head == nil
}

func (l *LinkedList) Head() *Node {
	return l.head
}

func (l *LinkedList) Index(value int) int {
	for i, n := range l.traverse() {
		if n.data == value {
			return i
		}
	}
	return -1
}

func (l *LinkedList) Pop() *Node {
	if l.head == nil {
		return nil
	}

	node := l.head
	l.head = l.head.next
	l.size--

	node.next = nil
	return node
}

func (l *LinkedList) PopIndex(idx int) *Node {
	if idx >= l.size {
		return nil
	} else if idx == 0 {
		return l.Pop()
	}
	nodes := l.traverse()
	for i, n := range nodes {
		if i == idx {
			l.size--
			if n.next == nil {
				nodes[i-1].next = nil
			} else {
				nodes[i-1].next = nodes[i+1]
			}
			n.next = nil
			return n
		}
	}
	return nil
}

func (l *LinkedList) Print() {
	for _, n := range l.traverse() {
		fmt.Println(n.data)
	}
}

func (l *LinkedList) Push(node *Node) {
	if l.Empty() {
		l.Add(node)
		return
	}
	tmp := l.head
	l.head = node
	node.next = tmp
	l.size++
}

func (l *LinkedList) PushIndex(node *Node, idx int) {
	if idx > l.size {
		return
	} else if l.size == 1 {
		l.Push(node)
		return
	} else if idx == 0 {
		l.Push(node)
		return
	}
	nodes := l.traverse()
	for i, n := range nodes {
		if i == idx {
			nodes[i-1].next = node
			node.next = n
			n.next = nodes[i+1]
			l.size++
			return
		}
	}

	// the only way we get here is if the idx is same as size
	nodes[len(nodes)-1].next = node
	node.next = nil
	l.size++
}

func (l *LinkedList) Reverse() {
	if l.size < 2 {
		// size of 0 is empty, size of 1 means its the same regardless, :-)
		return
	}
	var prev *Node
	current := l.Head()
	prev = nil
	i := 0
	for {
		next := current.next
		if i == 0 {
			current.next = nil
		} else {
			current.next = prev
		}
		prev = current
		current = next
		if current == nil {
			l.head = prev
			break
		}
		i++
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) traverse() []*Node {
	if l.size < 1 {
		return []*Node{}
	}

	r := make([]*Node, l.size)
	node := l.head

	i := 0
	for {
		if node.next == nil {
			r[i] = node
			return r
		}
		r[i] = node
		node = node.next
		i++
	}
}

func (l *LinkedList) Values() []int {
	vals := make([]int, l.size)
	for i, n := range l.traverse() {
		vals[i] = n.data
	}
	return vals
}

func (l *LinkedList) Apply(f func(i int) int) {
	for _, n := range l.traverse() {
			n.data = f(n.data)
	}
}
