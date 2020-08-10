package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

func NewNode(d int) *Node {
	return &Node{
		data: d,
		next: nil,
	}
}

type LL struct {
	head *Node
}

func NewLinkedList() *LL {
	return &LL{
		head: nil,
	}
}

func (l *LL) Add(n *Node) {
	if l.head == nil {
		l.head = n
		n.next = nil
		return
	}

	// linked list isnt empty, add to the "tail"
	node := l.head
	for {
		if node.next != nil {
			node = node.next
		} else {
			node.next = n
			return
		}
	}
}

func (ll *LL) Print() {
	node := ll.head
	for {
		if node.next == nil {
			return
		}
		fmt.Println(node.data)
		node = node.next
	}
}

func (l *LL) Remove(value int) {
	node := l.head
	if node == nil {
		return
	}
	previous := node
	for {
		if node == nil {
			return
		}
		if node.data == value {
			previous.next = node.next
			return
		}
		previous = node
		node = node.next
	}
}

func (l *LL) Traverse() []*Node {
	var r []*Node
	node := l.head
	for {
		if node.next == nil {
			r = append(r, node)
			return r
		}
		r = append(r, node)
	}
}

func (l *LL) Values() []int {
	vals := l.Traverse()
	r := make([]int, 0, len(vals))
	for _, v := range vals {
		r = append(r, v.data)
	}
	return r
}
