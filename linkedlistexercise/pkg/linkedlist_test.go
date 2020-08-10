package linkedlist

import (
	"fmt"
	"testing"
)

func TestCreateList(t *testing.T) {
	ll := NewLinkedList()
	for i := 0; i < 10; i++ {
		node := NewNode(i)
		ll.Add(node)
	}
	ll.Print()
	ll.Remove(6)
	ll.Print()
	fmt.Printf("vals: %+v\n", ll.Values())
	ll.Remove(6)
	ll.Print()
	fmt.Println("removed 0")
	ll.Remove(0)
	ll.Print()
}
